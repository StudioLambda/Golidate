package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/format"
	"github.com/studiolambda/golidate/rule"
	"github.com/studiolambda/golidate/translate/language"
)

func TestResultsTranslate(t *testing.T) {
	t.Run("Simple", func(t *testing.T) {
		results := golidate.Validate(
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
}

type NestedResults struct {
	Name    string
	Numbers []int
}

func (n NestedResults) Validate() golidate.Results {
	return golidate.Validate(
		golidate.Value(n.Name).Name("name").Rules(
			rule.MinLen(4),
		),
		golidate.Value(n.Numbers).Name("numbers").Rules(
			rule.SliceEach[int](
				rule.Min(1),
				rule.Max(10),
			),
		),
	)
}

func TestResultsGroup(t *testing.T) {
	nested := NestedResults{Numbers: []int{1, 2, 30}}

	results := golidate.Validate(
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

	require.Len(t, grouped, 4)

	_, ok := grouped["username"]
	_, ok2 := grouped["password"]

	require.True(t, ok)
	require.True(t, ok2)
}
