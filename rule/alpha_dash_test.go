package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestAlphaDash(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.AlphaDash()("foo_Bar-Baz")

		require.True(t, result.Passes())
		require.Equal(t, "alpha_dash", result.Code)
		require.Equal(t, "foo_Bar-Baz", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaDashRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.AlphaDash()("foo.bar@baz")

		require.False(t, result.Passes())
		require.Equal(t, "alpha_dash", result.Code)
		require.Equal(t, "foo.bar@baz", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaDashRegex.String()}, result.Metadata)
	})
}
