package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestMin(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Min(4)(5)

		require.True(t, result.Passes())
		require.Equal(t, "min", result.Code)
		require.Equal(t, 5, result.Value)
		require.Equal(t, golidate.Metadata{"min": int64(4)}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Min(4)(3)

		require.False(t, result.Passes())
		require.Equal(t, "min", result.Code)
		require.Equal(t, 3, result.Value)
		require.Equal(t, golidate.Metadata{"min": int64(4)}, result.Metadata)
	})

	t.Run("Exact", func(t *testing.T) {
		result := rule.Min(4)(4)

		require.True(t, result.Passes())
		require.Equal(t, "min", result.Code)
		require.Equal(t, 4, result.Value)
		require.Equal(t, golidate.Metadata{"min": int64(4)}, result.Metadata)
	})
}
