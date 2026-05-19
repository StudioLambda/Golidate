package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
)

var (
	// FailRule is a reusable test rule that always fails.
	FailRule = golidate.Rule(func(value any) golidate.Result {
		return golidate.Fail(value, "default.fail")
	})

	// PassRule is a reusable test rule that always passes.
	PassRule = golidate.Rule(func(value any) golidate.Result {
		return golidate.Pass(value, "default.pass")
	})
)

// TestRuleWith verifies Rule.With merges metadata into wrapped results.
func TestRuleWith(t *testing.T) {
	ruleWith := PassRule.
		With("key", "value").
		With("key2", "value")

	result := ruleWith(nil)

	require.Equal(t, golidate.Metadata{"key": "value", "key2": "value"}, result.Metadata)
}

// TestRuleWithMetadata verifies Rule.WithMetadata replaces metadata.
func TestRuleWithMetadata(t *testing.T) {
	ruleWithMetadata := PassRule.WithMetadata(golidate.Metadata{"key": "value"})
	result := ruleWithMetadata(nil)

	require.Equal(t, golidate.Metadata{"key": "value"}, result.Metadata)
}

// TestRuleWithMetadataMerged verifies Rule.WithMetadataMerged preserves old keys.
func TestRuleWithMetadataMerged(t *testing.T) {
	ruleWithMetadata := PassRule.
		WithMetadataMerged(golidate.Metadata{"key": "value"}).
		WithMetadataMerged(golidate.Metadata{"key2": "value"})

	result := ruleWithMetadata(nil)

	require.Equal(t, golidate.Metadata{"key": "value", "key2": "value"}, result.Metadata)
}

// TestRuleWithCode verifies Rule.Code changes only the result code.
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

// TestRuleCondition verifies false conditions make a result non-applicable.
func TestRuleCondition(t *testing.T) {
	rule := FailRule.Conditions(func(value any) bool {
		return value == "foo"
	})
	result := rule(nil)

	require.True(t, result.Passes())

	result = rule("foo")

	require.False(t, result.Passes())
}

// TestRuleWhen verifies boolean applicability conditions on rule results.
func TestRuleWhen(t *testing.T) {
	rule := FailRule.When(false)
	result := rule(nil)

	require.True(t, result.Passes())

	rule = FailRule.When(true)
	result = rule(nil)

	require.False(t, result.Passes())
}
