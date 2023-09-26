package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate/rule"
)

func TestNil(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Nil()(nil)

		require.True(t, result.Passes())
		require.Equal(t, "nil", result.Code)
		require.Equal(t, nil, result.Value)
		require.Empty(t, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Nil()(0)

		require.False(t, result.Passes())
		require.Equal(t, "nil", result.Code)
		require.Equal(t, 0, result.Value)
		require.Empty(t, result.Metadata)
	})
}
