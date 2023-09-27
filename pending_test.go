package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
	"github.com/studiolambda/golidate/translate/language"
)

func TestValue(t *testing.T) {
	value := golidate.Value(42)

	require.NotEmpty(t, value)
	require.IsType(t, golidate.Pending{}, value)
}

func TestPendingRules(t *testing.T) {
	result := golidate.
		Value(20).
		Rules(rule.Min(25)).
		Validate()

	require.False(t, result.PassesAll())
}

func TestPendingValidate(t *testing.T) {
	result := golidate.
		Value(20).
		Validate()

	require.True(t, result.PassesAll())
}

type Bs struct {
	Cp int
}

func (b Bs) Validate() golidate.Results {
	return golidate.Validate(
		golidate.Value(b.Cp).Name("cp").Rules(
			rule.Min(0),
			rule.Max(10),
		),
	)
}

type As struct {
	Bp Bs
	Dp []Bs
	Ep map[string]Bs
}

func (a As) Validate() golidate.Results {
	return golidate.Validate(
		golidate.Value(a.Bp).Name("bp"),
		golidate.Value(a.Dp).Name("dp"),
		golidate.Value(a.Ep).Name("ep"),
	)
}

func TestRecursiveValidate(t *testing.T) {
	a := As{
		Bp: Bs{Cp: 40},
		Dp: []Bs{
			{Cp: 10},
			{Cp: 21},
		},
		Ep: map[string]Bs{
			"foo": {Cp: 30},
			"bar": {Cp: 10},
		},
	}

	failed := golidate.
		Value(a).
		Validate().
		Failed().
		Translate(language.English)

	require.True(t, failed.Has("bp.cp"))
	require.True(t, failed.Has("dp.1.cp"))
	require.True(t, failed.Has("ep.foo.cp"))
}

func TestPointerValue(t *testing.T) {
	t.Run("PointerPass", func(t *testing.T) {
		realValue := 10
		value := golidate.Value(&realValue).Rules(
			rule.Min(0),
			rule.Max(20),
		)
		result := value.Validate()

		require.True(t, result.PassesAll())
	})

	t.Run("PointerFail", func(t *testing.T) {
		realValue := 30
		value := golidate.Value(&realValue).Rules(
			rule.Min(0),
			rule.Max(20),
		)
		result := value.Validate()

		require.False(t, result.PassesAll())
	})

	t.Run("Nil", func(t *testing.T) {
		var realValue *int
		value := golidate.Value(realValue).Rules(
			rule.Min(0),
			rule.Max(20),
		)
		result := value.Validate()

		require.False(t, result.PassesAll())
	})

	t.Run("NilChecked", func(t *testing.T) {
		value := golidate.Value(nil).Rules(
			rule.Or(rule.Nil(), rule.Min(0)),
			rule.Or(rule.Nil(), rule.Max(20)),
		)
		result := value.Validate()

		require.True(t, result.PassesAll())

		var realValue2 *int = nil
		value2 := golidate.Value(realValue2).Rules(
			rule.Or(rule.Nil(), rule.Min(0)),
			rule.Or(rule.Nil(), rule.Max(20)),
		)
		result2 := value2.Validate()

		require.True(t, result2.PassesAll())

		realValue3 := 10
		value3 := golidate.Value(&realValue3).Rules(
			rule.Or(rule.Nil(), rule.Min(0)),
			rule.Or(rule.Nil(), rule.Max(20)),
		)
		result3 := value3.Validate()

		require.True(t, result3.PassesAll())
	})
}
