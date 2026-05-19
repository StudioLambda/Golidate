package golidate_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
	"github.com/studiolambda/golidate/translate/language"
)

// Self is a test value that validates itself through golidate.Self.
type Self int

// Validate applies numeric rules to the underlying Self integer.
func (s Self) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(ctx, golidate.Self(int(s)).Name("self").Rules(
		rule.Min(0), rule.Max(10),
	))
}

// PointerOnly is a test validator with only a pointer receiver Validate method.
type PointerOnly struct {
	// Cp stores the number checked by pointer-receiver validation.
	Cp int
}

// Validate applies numeric rules through a pointer receiver.
func (p *PointerOnly) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(p.Cp).Name("cp").Rules(
			rule.Min(0),
			rule.Max(10),
		),
	)
}

// TestValue verifies Value returns a non-empty Pending builder.
func TestValue(t *testing.T) {
	value := golidate.Value(42)

	require.NotEmpty(t, value)
	require.IsType(t, golidate.Pending{}, value)
}

// TestSelf verifies Self disables recursive Validator handling for a value.
func TestSelf(t *testing.T) {
	results := Self(5).Validate(context.Background())

	require.True(t, results.PassesAll())
}

// TestPendingRules verifies pending rules are applied to the value.
func TestPendingRules(t *testing.T) {
	result := golidate.
		Value(20).
		Rules(rule.Min(25)).
		Validate(context.Background())

	require.False(t, result.PassesAll())
}

// TestPendingValidate verifies a pending value with no rules passes.
func TestPendingValidate(t *testing.T) {
	result := golidate.
		Value(20).
		Validate(context.Background())

	require.True(t, result.PassesAll())
}

// Bs is a nested validator used by recursive pending validation tests.
type Bs struct {
	// Cp stores the number checked by Bs validation.
	Cp int
}

// Validate applies numeric rules to Bs.Cp.
func (b Bs) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(b.Cp).Name("cp").Rules(
			rule.Min(0),
			rule.Max(10),
		),
	)
}

// As is a parent validator containing nested struct, slice, and map values.
type As struct {
	// Bp stores a direct nested validator.
	Bp Bs
	// Dp stores nested validators in a slice.
	Dp []Bs
	// Ep stores nested validators in a map.
	Ep map[string]Bs
}

// Validate delegates to the nested validators stored in As.
func (a As) Validate(ctx context.Context) golidate.Results {
	return golidate.Validate(
		ctx,
		golidate.Value(a.Bp).Name("bp"),
		golidate.Value(a.Dp).Name("dp"),
		golidate.Value(a.Ep).Name("ep"),
	)
}

// TestRecursiveValidate verifies recursive validation prefixes nested attributes.
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
		Validate(context.Background()).
		Failed().
		Translate(language.English)

	require.True(t, failed.Has("bp.cp"))
	require.True(t, failed.Has("dp.1.cp"))
	require.True(t, failed.Has("ep.foo.cp"))
}

// TestRecursiveValidateFormatsNonStringMapKeys verifies formatted key attributes.
func TestRecursiveValidateFormatsNonStringMapKeys(t *testing.T) {
	values := map[int]Bs{
		42: {Cp: 30},
	}

	failed := golidate.
		Value(values).
		Name("items").
		Validate(context.Background()).
		Failed()

	require.True(t, failed.Has("items.42.cp"))
}

// TestRecursiveValidateMapOrder verifies deterministic recursive map ordering.
func TestRecursiveValidateMapOrder(t *testing.T) {
	values := map[string]Bs{
		"c": {Cp: 30},
		"a": {Cp: 30},
		"b": {Cp: 30},
	}

	failed := golidate.
		Value(values).
		Name("items").
		Validate(context.Background()).
		Failed()

	require.Len(t, failed, 3)
	require.Equal(t, "items.a.cp", failed[0].Attribute)
	require.Equal(t, "items.b.cp", failed[1].Attribute)
	require.Equal(t, "items.c.cp", failed[2].Attribute)
}

// TestRecursiveValidatePointerReceiver verifies pointer-only validators run.
func TestRecursiveValidatePointerReceiver(t *testing.T) {
	value := &PointerOnly{Cp: 40}

	failed := golidate.
		Value(value).
		Validate(context.Background()).
		Failed().
		Translate(language.English)

	require.True(t, failed.Has("cp"))
}

// TestPointerValue verifies pointer dereferencing and nil pointer behavior.
func TestPointerValue(t *testing.T) {
	t.Run("PointerPass", func(t *testing.T) {
		realValue := 10
		value := golidate.Value(&realValue).Rules(
			rule.Min(0),
			rule.Max(20),
		)
		result := value.Validate(context.Background())

		require.True(t, result.PassesAll())
	})

	t.Run("PointerFail", func(t *testing.T) {
		realValue := 30
		value := golidate.Value(&realValue).Rules(
			rule.Min(0),
			rule.Max(20),
		)
		result := value.Validate(context.Background())

		require.False(t, result.PassesAll())
	})

	t.Run("Nil", func(t *testing.T) {
		var realValue *int
		value := golidate.Value(realValue).Rules(
			rule.Min(0),
			rule.Max(20),
		)
		result := value.Validate(context.Background())

		require.False(t, result.PassesAll())
	})

	t.Run("NilChecked", func(t *testing.T) {
		value := golidate.Value(nil).Rules(
			rule.Optional(rule.Min(0)),
			rule.Optional(rule.Max(20)),
		)
		result := value.Validate(context.Background())

		require.True(t, result.PassesAll())

		var realValue2 *int = nil
		value2 := golidate.Value(realValue2).Rules(
			rule.Optional(rule.Min(0)),
		)
		result2 := value2.Validate(context.Background())

		require.True(t, result2.PassesAll())

		realValue3 := 10
		value3 := golidate.Value(&realValue3).Rules(
			rule.Optional(rule.Min(0)),
			rule.Optional(rule.Max(20)),
		)
		result3 := value3.Validate(context.Background())

		require.True(t, result3.PassesAll())
	})
}
