package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

// TestAccepted verifies accepted string values pass and other values fail.
func TestAccepted(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		for _, value := range rule.AcceptedValues {
			result := rule.Accepted()(value)
			require.True(t, result.Passes())
			require.Equal(t, "accepted", result.Code)
			require.Equal(t, value, result.Value)
			require.Equal(t, golidate.Metadata{"values": rule.AcceptedValues}, result.Metadata)
		}
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Accepted()("something")

		require.False(t, result.Passes())
		require.Equal(t, "accepted", result.Code)
		require.Equal(t, "something", result.Value)
		require.Equal(t, golidate.Metadata{"values": rule.AcceptedValues}, result.Metadata)
	})
}

// TestAcceptedUsesCopiedValues verifies Accepted does not share exported slice state.
func TestAcceptedUsesCopiedValues(t *testing.T) {
	rule.AcceptedValues[0] = "changed"
	t.Cleanup(func() {
		rule.AcceptedValues[0] = "true"
	})
	accepted := rule.Accepted()

	result := accepted("true")

	require.True(t, result.Passes())
}
