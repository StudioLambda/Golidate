package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestTypeOf(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result := rule.TypeOf("")("foo")

		require.True(t, result.Passes())
		require.Equal(t, "type_of", result.Code)
		require.Equal(t, "foo", result.Value)
		require.Equal(t, golidate.Metadata{"type": "string"}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.TypeOf("")(0)

		require.False(t, result.Passes())
		require.Equal(t, "type_of", result.Code)
		require.Equal(t, 0, result.Value)
		require.Equal(t, golidate.Metadata{"type": "string"}, result.Metadata)
	})
}
