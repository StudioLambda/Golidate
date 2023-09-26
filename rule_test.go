package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
)

var (
	FailRule = golidate.Rule(func(value any) golidate.Result {
		return golidate.Fail(value, "default.fail")
	})

	PassRule = golidate.Rule(func(value any) golidate.Result {
		return golidate.Pass(value, "default.pass")
	})
)

func TestRuleWith(t *testing.T) {
	ruleWith := PassRule.
		With("key", "value").
		With("key2", "value")

	result := ruleWith(nil)

	require.Equal(t, golidate.Metadata{"key": "value", "key2": "value"}, result.Metadata)
}

func TestRuleWithMetadata(t *testing.T) {
	ruleWithMetadata := PassRule.WithMetadata(golidate.Metadata{"key": "value"})
	result := ruleWithMetadata(nil)

	require.Equal(t, golidate.Metadata{"key": "value"}, result.Metadata)
}

func TestRuleWithMetadataMerged(t *testing.T) {
	ruleWithMetadata := PassRule.
		WithMetadataMerged(golidate.Metadata{"key": "value"}).
		WithMetadataMerged(golidate.Metadata{"key2": "value"})

	result := ruleWithMetadata(nil)

	require.Equal(t, golidate.Metadata{"key": "value", "key2": "value"}, result.Metadata)
}

func TestRuleWithCode(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := PassRule.Code("custom")(nil)

		require.True(t, result.Passes())
		require.Equal(t, "custom", result.Code)
	})

	t.Run("Failure", func(t *testing.T) {
		result := FailRule.Code("custom")(nil)

		require.False(t, result.Passes())
		require.Equal(t, "custom", result.Code)
	})
}

func TestRuleCondition(t *testing.T) {
	rule := FailRule.Conditions(func(value any) bool {
		return value == "foo"
	})
	result := rule(nil)

	require.True(t, result.Passes())

	result = rule("foo")

	require.False(t, result.Passes())
}

func TestRuleWhen(t *testing.T) {
	rule := FailRule.When(false)
	result := rule(nil)

	require.True(t, result.Passes())

	rule = FailRule.When(true)
	result = rule(nil)

	require.False(t, result.Passes())
}
