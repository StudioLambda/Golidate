package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestHasPrefix(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.HasPrefix("foo")("foo bar")

		require.True(t, result.Passes())
		require.Equal(t, "has_prefix", result.Code)
		require.Equal(t, "foo bar", result.Value)
		require.Equal(t, golidate.Metadata{"prefix": "foo"}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.HasPrefix("foo")("bar baz")

		require.False(t, result.Passes())
		require.Equal(t, "has_prefix", result.Code)
		require.Equal(t, "bar baz", result.Value)
		require.Equal(t, golidate.Metadata{"prefix": "foo"}, result.Metadata)
	})
}
