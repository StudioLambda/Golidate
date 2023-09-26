package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestAlpha(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Alpha()("fooBarBaz")

		require.True(t, result.Passes())
		require.Equal(t, "alpha", result.Code)
		require.Equal(t, "fooBarBaz", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Alpha()("foo.bar@baz")

		require.False(t, result.Passes())
		require.Equal(t, "alpha", result.Code)
		require.Equal(t, "foo.bar@baz", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaRegex.String()}, result.Metadata)
	})
}
