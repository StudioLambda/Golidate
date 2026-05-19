package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

// TestMaxLen verifies maximum length validation behavior.
func TestMaxLen(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.MaxLen(4)("asd")

		require.True(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, "asd", result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("ArrayPass", func(t *testing.T) {
		value := [3]int{1, 2, 3}
		result := rule.MaxLen(4)(value)

		require.True(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("SlicePass", func(t *testing.T) {
		value := []int{1, 2, 3}
		result := rule.MaxLen(4)(value)

		require.True(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("MapPass", func(t *testing.T) {
		value := map[string]int{"a": 1, "b": 2, "c": 3}
		result := rule.MaxLen(4)(value)

		require.True(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("ChanPass", func(t *testing.T) {
		value := make(chan int, 3)
		value <- 1
		value <- 2
		value <- 3
		result := rule.MaxLen(4)(value)

		require.True(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.MaxLen(4)("asdfg")

		require.False(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, "asdfg", result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("Exact", func(t *testing.T) {
		result := rule.MaxLen(4)("asdf")

		require.True(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, "asdf", result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("Unsupported", func(t *testing.T) {
		result := rule.MaxLen(4)(10)

		require.False(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, 10, result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})

	t.Run("Nil", func(t *testing.T) {
		result := rule.MaxLen(4)(nil)

		require.False(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, nil, result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})
}
