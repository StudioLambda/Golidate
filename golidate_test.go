package golidate_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

// A is a marker type used by polymorphic validation tests.
type A struct{}

// B is a marker type used by polymorphic validation tests.
type B struct{}

// EitherAorB validates its value against one of two possible marker types.
type EitherAorB struct {
	// kind selects which marker type should be accepted.
	kind string
	// value stores the marker instance being validated.
	value any
}

// Validate applies the conditional type rule selected by kind.
func (either EitherAorB) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(ctx).Values(
		golidate.Value(either.value).Rules(
			rule.Type[A]().When(either.kind == "a"),
			rule.Type[B]().When(either.kind == "b"),
		),
	)
}

// Profile is a nested test validator for user profile fields.
type Profile struct {
	// Name stores the profile display name used in nested validation.
	Name string
	// Email stores the profile email used in nested validation.
	Email string
	// Age stores the profile age checked by a minimum rule.
	Age int
}

// User is a parent test validator containing direct and nested values.
type User struct {
	// Username stores the user name checked by a non-nil rule.
	Username string
	// Password stores the password checked by a non-nil rule.
	Password string
	// Profile stores nested profile data validated recursively.
	Profile Profile
}

// Validate validates direct user fields and delegates to Profile validation.
func (user User) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(ctx).Values(
		golidate.Value(user.Username).Rules(
			rule.Not(rule.Nil()),
		),
		golidate.Value(user.Password).Rules(
			rule.Not(rule.Nil()),
		),
		golidate.Value(user.Profile),
	)
}

// Validate validates profile fields used by nested validation tests.
func (user Profile) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(ctx).Values(
		golidate.Value(user.Name),
		golidate.Value(user.Email).Rules(
			rule.Not(rule.Nil()),
		),
		golidate.Value(user.Age).Rules(
			rule.Min(18),
		),
	)
}

// TestValidate verifies top-level and nested validation pass/fail behavior.
func TestValidate(t *testing.T) {
	t.Run("SinglePass", func(t *testing.T) {
		results := golidate.Validate(context.Background()).Values(
			golidate.Value(3).Rules(
				rule.Min(2),
			),
		)

		require.True(t, results.PassesAll())
	})

	t.Run("SingleFailure", func(t *testing.T) {
		result := golidate.Validate(context.Background()).Values(
			golidate.Value(1).Rules(
				rule.Min(2),
			),
		)

		require.False(t, result.PassesAll())
	})

	t.Run("SingleMultipleFailures", func(t *testing.T) {
		result := golidate.Validate(context.Background()).Values(
			golidate.Value(1).Rules(rule.Min(2), rule.Max(0)),
		)

		require.False(t, result.PassesAll())
	})

	t.Run("MultipleSuccess", func(t *testing.T) {
		result := golidate.Validate(context.Background()).Values(
			golidate.Value(3).Rules(rule.Min(2)),
			golidate.Value(3).Rules(rule.Max(4)),
		)

		require.True(t, result.PassesAll())
	})

	t.Run("MultipleSingleFailure", func(t *testing.T) {
		result := golidate.Validate(context.Background()).Values(
			golidate.Value(1).Rules(rule.Min(2)),
			golidate.Value(3).Rules(rule.Max(4)),
		)

		require.False(t, result.PassesAll())
	})

	t.Run("MultipleMultipleFailures", func(t *testing.T) {
		result := golidate.Validate(context.Background()).Values(
			golidate.Value(1).Rules(rule.Min(2)),
			golidate.Value(5).Rules(rule.Max(4)),
		)

		require.False(t, result.PassesAll())
	})

	t.Run("NestedSuccess", func(t *testing.T) {
		user := User{
			Username: "john.doe",
			Password: "secret",
			Profile: Profile{
				Name:  "Jhon Doe",
				Email: "john.doe@example.com",
				Age:   35,
			},
		}

		result := golidate.Validate(context.Background()).Values(
			golidate.Value(user),
		)

		require.True(t, result.PassesAll())
	})

	t.Run("NestedFailure", func(t *testing.T) {
		user := User{
			Username: "john.doe",
			Password: "secret",
			Profile: Profile{
				Name:  "Jhon Doe",
				Email: "john.doe@example.com",
				Age:   16,
			},
		}

		result := golidate.Validate(context.Background()).Values(
			golidate.Value(user),
		)

		require.False(t, result.PassesAll())
	})
}

// TestPolymorphism verifies conditionally applied type rules.
func TestPolymorphism(t *testing.T) {
	t.Run("SuccessA", func(t *testing.T) {
		result := EitherAorB{kind: "a", value: A{}}.Validate(context.Background())

		require.True(t, result.PassesAll())
	})

	t.Run("SuccessB", func(t *testing.T) {
		result := EitherAorB{kind: "b", value: B{}}.Validate(context.Background())

		require.True(t, result.PassesAll())
	})

	t.Run("Failure", func(t *testing.T) {
		result := EitherAorB{kind: "a", value: B{}}.Validate(context.Background())

		require.False(t, result.PassesAll())
	})
}

// ContextValue validates a number using a maximum read from context.
type ContextValue struct {
	// Value stores the number checked by context-driven limits.
	Value int
}

// maxKeyStruct is the private context key type for maximum values.
type maxKeyStruct struct{}

// MaxKey is the context key used by ContextValue validation.
var MaxKey = maxKeyStruct{}

// Validate validates ContextValue using a context-provided maximum.
func (c ContextValue) Validate(ctx context.Context) golidate.Results {
	max := ctx.Value(MaxKey).(int64)
	return golidate.Validate(ctx).Values(
		golidate.Value(c.Value).Rules(
			rule.Min(0),
			rule.Max(max),
		),
	)
}

// TestContextValues verifies validators can consume context values.
func TestContextValues(t *testing.T) {
	value := ContextValue{Value: 10}
	ctx := context.WithValue(context.Background(), MaxKey, int64(15))
	result := golidate.Value(value).Validate(ctx)

	require.True(t, result.PassesAll())

	ctx = context.WithValue(context.Background(), MaxKey, int64(5))
	result = golidate.Value(value).Validate(ctx)

	require.False(t, result.PassesAll())
}

// TestComplexErrorKeys verifies nested slice and map attribute naming.
func TestComplexErrorKeys(t *testing.T) {
	t.Run("Slices", func(t *testing.T) {
		value := [][]int{{1, 5}, {10, 15}, {25, 30}}
		result := golidate.Value(value).Name("foo").Rules(
			rule.SliceValues[[][]int](
				rule.SliceValues[[]int](
					rule.Min(10),
					rule.Max(15),
				),
			),
		).Validate(context.Background())

		failed := result.Failed()

		require.True(t, failed.Has("foo.0.0"))
		require.True(t, failed.Has("foo.0.1"))
		require.True(t, failed.Has("foo.2.0"))
		require.True(t, failed.Has("foo.2.1"))
	})

	t.Run("Maps", func(t *testing.T) {
		value := map[string]int{
			"one":   5,
			"two":   10,
			"three": 25,
		}

		result := golidate.Value(value).Name("foo").Rules(
			rule.MapValues[map[string]int](
				rule.Min(10),
				rule.Max(15),
			),
		).Validate(context.Background())

		failed := result.Failed()

		require.True(t, failed.Has("foo.one"))
		require.True(t, failed.Has("foo.three"))
	})

	t.Run("Map Keys", func(t *testing.T) {
		value := map[string]int{
			"one":     5,
			"two bad": 10,
			"three":   25,
		}

		result := golidate.Value(value).Name("foo").Rules(
			rule.MapKeys[map[string]int](
				rule.Alpha(),
			),
		).Validate(context.Background())

		failed := result.Failed()

		require.True(t, failed.Has("foo.[two bad]"))
	})
}
