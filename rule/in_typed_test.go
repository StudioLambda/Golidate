package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestInTyped(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		values := []string{"a", "b", "c"}
		result := rule.InTyped(values...)("b")

		require.True(t, result.Passes())
		require.Equal(t, "in_typed", result.Code)
		require.Equal(t, "b", result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		values := []string{"a", "b", "c"}
		result := rule.InTyped(values...)("d")

		require.False(t, result.Passes())
		require.Equal(t, "in_typed", result.Code)
		require.Equal(t, "d", result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})

	t.Run("InvalidValue", func(t *testing.T) {
		values := []string{"a", "b", "c"}
		value := map[string]string{}
		result := rule.InTyped(values...)(value)

		require.False(t, result.Passes())
		require.Equal(t, "in_typed", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})

	t.Run("InvalidValues", func(t *testing.T) {
		values := []string{"a", "b", "c"}
		result := rule.InTyped(values...)("d")

		require.False(t, result.Passes())
		require.Equal(t, "in_typed", result.Code)
		require.Equal(t, "d", result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})
}
