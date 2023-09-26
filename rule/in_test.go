package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestIn(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		values := []any{"a", "b", "c"}
		result := rule.In(values...)("b")

		require.True(t, result.Passes())
		require.Equal(t, "in", result.Code)
		require.Equal(t, "b", result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		values := []any{"a", "b", "c"}
		result := rule.In(values...)("d")

		require.False(t, result.Passes())
		require.Equal(t, "in", result.Code)
		require.Equal(t, "d", result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})

	t.Run("InvalidValue", func(t *testing.T) {
		values := []any{"a", "b", "c"}
		value := map[string]string{}
		result := rule.In(values...)(value)

		require.False(t, result.Passes())
		require.Equal(t, "in", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})

	t.Run("InvalidValues", func(t *testing.T) {
		values := []any{"a", map[string]string{}, "c"}
		result := rule.In(values...)("b")

		require.False(t, result.Passes())
		require.Equal(t, "in", result.Code)
		require.Equal(t, "b", result.Value)
		require.Equal(t, golidate.Metadata{"values": values}, result.Metadata)
	})
}
