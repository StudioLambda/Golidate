package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestMaxLen(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.MaxLen(4)("asd")

		require.True(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, "asd", result.Value)
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

	t.Run("Panic", func(t *testing.T) {
		result := rule.MaxLen(4)(10)

		require.False(t, result.Passes())
		require.Equal(t, "max_len", result.Code)
		require.Equal(t, 10, result.Value)
		require.Equal(t, golidate.Metadata{"max": 4}, result.Metadata)
	})
}
