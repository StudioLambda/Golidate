package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestLen(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Len(4)("asdf")

		require.True(t, result.Passes())
		require.Equal(t, "len", result.Code)
		require.Equal(t, "asdf", result.Value)
		require.Equal(t, golidate.Metadata{"len": 4}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Len(4)("asd")

		require.False(t, result.Passes())
		require.Equal(t, "len", result.Code)
		require.Equal(t, "asd", result.Value)
		require.Equal(t, golidate.Metadata{"len": 4}, result.Metadata)
	})

	t.Run("Panic", func(t *testing.T) {
		result := rule.Len(4)(nil)

		require.False(t, result.Passes())
		require.Equal(t, "len", result.Code)
		require.Equal(t, nil, result.Value)
		require.Equal(t, golidate.Metadata{"len": 4}, result.Metadata)
	})
}
