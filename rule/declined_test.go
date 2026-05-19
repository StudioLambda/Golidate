package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

// TestDeclined verifies declined string values pass and other values fail.
func TestDeclined(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		for _, value := range rule.DeclinedValues {
			result := rule.Declined()(value)

			require.True(t, result.Passes())
			require.Equal(t, "declined", result.Code)
			require.Equal(t, value, result.Value)
			require.Equal(t, golidate.Metadata{"values": rule.DeclinedValues}, result.Metadata)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Declined()("something")

		require.False(t, result.Passes())
		require.Equal(t, "declined", result.Code)
		require.Equal(t, "something", result.Value)
		require.Equal(t, golidate.Metadata{"values": rule.DeclinedValues}, result.Metadata)
	})
}

// TestDeclinedUsesCopiedValues verifies Declined does not share exported slice state.
func TestDeclinedUsesCopiedValues(t *testing.T) {
	rule.DeclinedValues[0] = "changed"
	t.Cleanup(func() {
		rule.DeclinedValues[0] = "false"
	})
	declined := rule.Declined()

	result := declined("false")

	require.True(t, result.Passes())
}
