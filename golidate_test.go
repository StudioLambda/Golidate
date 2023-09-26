package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

type A struct{}
type B struct{}
type EitherAorB struct {
	kind  string
	value any
}

func (either EitherAorB) Validate() golidate.Results {
	return golidate.Validate(
		golidate.Value(either.value).Rules(
			rule.Type[A]().When(either.kind == "a"),
			rule.Type[B]().When(either.kind == "b"),
		),
	)
}

type Profile struct {
	Name  string
	Email string
	Age   int
}

type User struct {
	Username string
	Password string
	Profile  Profile
}

func (user User) Validate() golidate.Results {
	return golidate.Validate(
		golidate.Value(user.Username).Rules(
			rule.Not(rule.Nil()),
		),
		golidate.Value(user.Password).Rules(
			rule.Not(rule.Nil()),
		),
		golidate.Value(user.Profile),
	)
}

func (user Profile) Validate() golidate.Results {
	return golidate.Validate(
		golidate.Value(user.Name),
		golidate.Value(user.Email).Rules(
			rule.Not(rule.Nil()),
		),
		golidate.Value(user.Age).Rules(
			rule.Min(18),
		),
	)
}

func TestValidate(t *testing.T) {
	t.Run("SinglePass", func(t *testing.T) {
		results := golidate.Validate(
			golidate.Value(3).Rules(
				rule.Min(2),
			),
		)

		require.True(t, results.PassesAll())
	})

	t.Run("SingleFailure", func(t *testing.T) {
		result := golidate.Validate(
			golidate.Value(1).Rules(
				rule.Min(2),
			),
		)

		require.False(t, result.PassesAll())
	})

	t.Run("SingleMultipleFailures", func(t *testing.T) {
		result := golidate.Validate(
			golidate.Value(1).Rules(rule.Min(2), rule.Max(0)),
		)

		require.False(t, result.PassesAll())
	})

	t.Run("MultipleSuccess", func(t *testing.T) {
		result := golidate.Validate(
			golidate.Value(3).Rules(rule.Min(2)),
			golidate.Value(3).Rules(rule.Max(4)),
		)

		require.True(t, result.PassesAll())
	})

	t.Run("MultipleSingleFailure", func(t *testing.T) {
		result := golidate.Validate(
			golidate.Value(1).Rules(rule.Min(2)),
			golidate.Value(3).Rules(rule.Max(4)),
		)

		require.False(t, result.PassesAll())
	})

	t.Run("MultipleMultipleFailures", func(t *testing.T) {
		result := golidate.Validate(
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

		result := golidate.Validate(
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

		result := golidate.Validate(
			golidate.Value(user),
		)

		require.False(t, result.PassesAll())
	})
}

func TestPolymorphism(t *testing.T) {
	t.Run("SuccessA", func(t *testing.T) {
		result := EitherAorB{kind: "a", value: A{}}.Validate()

		require.True(t, result.PassesAll())
	})

	t.Run("SuccessB", func(t *testing.T) {
		result := EitherAorB{kind: "b", value: B{}}.Validate()

		require.True(t, result.PassesAll())
	})

	t.Run("Failure", func(t *testing.T) {
		result := EitherAorB{kind: "a", value: B{}}.Validate()

		require.False(t, result.PassesAll())
	})
}
