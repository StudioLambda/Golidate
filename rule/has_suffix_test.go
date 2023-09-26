package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestHasSuffix(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.HasSuffix("bar")("foo bar")

		require.True(t, result.Passes())
		require.Equal(t, "has_suffix", result.Code)
		require.Equal(t, "foo bar", result.Value)
		require.Equal(t, golidate.Metadata{"suffix": "bar"}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.HasSuffix("bar")("bar baz")

		require.False(t, result.Passes())
		require.Equal(t, "has_suffix", result.Code)
		require.Equal(t, "bar baz", result.Value)
		require.Equal(t, golidate.Metadata{"suffix": "bar"}, result.Metadata)
	})
}
