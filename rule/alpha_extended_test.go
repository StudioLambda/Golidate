package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestAlphaExtended(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.AlphaExtended()("123foo123.Bar_123Baz123")

		require.True(t, result.Passes())
		require.Equal(t, "alpha_extended", result.Code)
		require.Equal(t, "123foo123.Bar_123Baz123", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaExtendedRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.AlphaExtended()("foo.bar@baz")

		require.False(t, result.Passes())
		require.Equal(t, "alpha_extended", result.Code)
		require.Equal(t, "foo.bar@baz", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaExtendedRegex.String()}, result.Metadata)
	})
}
