package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestZero(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result := rule.Zero()(0)

		require.True(t, result.Passes())
		require.Equal(t, "zero", result.Code)
		require.Equal(t, 0, result.Value)
		require.Equal(t, golidate.Metadata{"zero": 0}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Zero()(10)

		require.False(t, result.Passes())
		require.Equal(t, "zero", result.Code)
		require.Equal(t, 10, result.Value)
		require.Equal(t, golidate.Metadata{"zero": 0}, result.Metadata)
	})
}
