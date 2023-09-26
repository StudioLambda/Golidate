package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestConvertible(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result := rule.Convertible[string]()(0)

		require.True(t, result.Passes())
		require.Equal(t, "convertible", result.Code)
		require.Equal(t, 0, result.Value)
		require.Equal(t, golidate.Metadata{"type": "string"}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Convertible[int]()("")

		require.False(t, result.Passes())
		require.Equal(t, "convertible", result.Code)
		require.Equal(t, "", result.Value)
		require.Equal(t, golidate.Metadata{"type": "int"}, result.Metadata)
	})
}
