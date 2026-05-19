package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestBoolean(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		for _, value := range rule.BooleanValues {
			result := rule.Boolean()(value)

			require.True(t, result.Passes())
			require.Equal(t, "boolean", result.Code)
			require.Equal(t, value, result.Value)
			require.Equal(t, golidate.Metadata{"values": rule.BooleanValues}, result.Metadata)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Boolean()("something")

		require.False(t, result.Passes())
		require.Equal(t, "boolean", result.Code)
		require.Equal(t, "something", result.Value)
		require.Equal(t, golidate.Metadata{"values": rule.BooleanValues}, result.Metadata)
	})
}

func TestBooleanUsesCopiedValues(t *testing.T) {
	rule.BooleanValues[0] = "changed"
	t.Cleanup(func() {
		rule.BooleanValues[0] = "true"
	})
	boolean := rule.Boolean()

	result := boolean("true")

	require.True(t, result.Passes())
}
