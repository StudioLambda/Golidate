package golidate_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/format"
	"github.com/studiolambda/golidate/rule"
	"github.com/studiolambda/golidate/translate/language"
)

// TestResultsTranslate verifies dictionary translation and override behavior.
func TestResultsTranslate(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		results := golidate.Validate(
			context.Background(),
			golidate.Value("something invalid").Name("username").Rules(
				rule.Alpha(),
			),
		)

		translated := results.Failed().Translate(language.English)

		require.Len(t, translated, 1)
		require.Equal(t, "the username field must only contain letters", translated[0].Message)
	})

	t.Run("Complex", func(t *testing.T) {
		results := golidate.Validate(
			context.Background(),
			golidate.Value("something invalid").Name("username").Rules(
				rule.And(rule.MinLen(5), rule.MaxLen(10)),
				rule.Not(rule.Nil()),
			),
		)

		require.Len(t, results, 2)

		translated := results.Failed().Translate(language.English)

		require.Len(t, translated, 1)

		expected := language.English["and"](language.English, translated[0])

		require.Equal(t, expected.Message, translated[0].Message)
	})

	t.Run("OverridesDefaultsForMultipleResults", func(t *testing.T) {
		results := golidate.Results{
			golidate.Fail("ab", "min_len").Name("username").With("min", 3),
			golidate.Fail("cd", "min_len").Name("password").With("min", 3),
		}
		overrides := golidate.Dictionary{
			"min_len": func(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
				result.Message = "override " + result.Attribute

				return result
			},
		}

		translated := results.Translate(language.English, overrides)

		require.Len(t, translated, 2)
		require.Equal(t, "override username", translated[0].Message)
		require.Equal(t, "override password", translated[1].Message)
	})

	t.Run("UsesDefaultsForMultipleResults", func(t *testing.T) {
		results := golidate.Results{
			golidate.Fail("ab", "min_len").Name("username").With("min", 3),
			golidate.Fail("cd", "min_len").Name("password").With("min", 3),
		}

		translated := results.Translate(language.English)

		require.Len(t, translated, 2)
		require.Equal(t, "the username field must be at least 3 characters long", translated[0].Message)
		require.Equal(t, "the password field must be at least 3 characters long", translated[1].Message)
	})
}

// NestedResults is a test validator with direct and nested child failures.
type NestedResults struct {
	// Name stores the value checked by a minimum length rule.
	Name string
	// Numbers stores values checked through SliceValues.
	Numbers []int
}

// Validate validates NestedResults fields for grouping tests.
func (n NestedResults) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(n.Name).Name("name").Rules(
			rule.MinLen(4),
		),
		golidate.Value(n.Numbers).Name("numbers").Rules(
			rule.SliceValues[[]int](
				rule.Min(1),
				rule.Max(10),
			),
		),
	)
}

// TestResultsGroup verifies failed translated results can be grouped by attribute.
func TestResultsGroup(t *testing.T) {
	nested := NestedResults{Numbers: []int{1, 2, 30}}

	results := golidate.Validate(
		context.Background(),
		golidate.Value("something.valid").Name("username").Rules(
			rule.Not(rule.Nil()),
			rule.Type[string](),
			rule.AlphaDash(),
			rule.MinLen(2),
			rule.MaxLen(255),
		),
		golidate.Value(nested).Name("nested"),
		golidate.Value("something invalid").Name("password").Rules(
			rule.Not(rule.Nil()),
			rule.Type[string](),
			rule.AlphaDash(),
			rule.MinLen(2),
			rule.MaxLen(255),
		),
	)

	grouped := results.
		Failed().
		Translate(language.English).
		Group().
		Messages(format.Capitalize(), format.Punctuate())

	require.Len(t, grouped, 5)

	_, ok := grouped["username"]
	_, ok2 := grouped["password"]

	require.True(t, ok)
	require.True(t, ok2)
}

// TestResultsClassifiesUnflattenedChildren verifies direct child state is included.
func TestResultsClassifiesUnflattenedChildren(t *testing.T) {
	results := golidate.Results{
		rule.SliceValues[[]int](rule.Max(10))([]int{5, 20}),
	}

	require.False(t, results.PassesAll())
	require.False(t, results.PassesAny())
	require.Len(t, results.Failed(), 1)
	require.Empty(t, results.Passed())
}
