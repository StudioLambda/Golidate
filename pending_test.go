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
