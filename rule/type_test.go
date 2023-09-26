package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestType(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result := rule.Type[string]()("")

		require.True(t, result.Passes())
		require.Equal(t, "type", result.Code)
		require.Equal(t, "", result.Value)
		require.Equal(t, golidate.Metadata{"type": "string"}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Type[string]()(0)

		require.False(t, result.Passes())
		require.Equal(t, "type", result.Code)
		require.Equal(t, 0, result.Value)
		require.Equal(t, golidate.Metadata{"type": "string"}, result.Metadata)
	})
}
