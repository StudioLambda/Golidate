package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestAlphaNumeric(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.AlphaNumeric()("123foo123Bar123Baz123")

		require.True(t, result.Passes())
		require.Equal(t, "alpha_numeric", result.Code)
		require.Equal(t, "123foo123Bar123Baz123", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaNumericRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.AlphaNumeric()("foo.bar@baz")

		require.False(t, result.Passes())
		require.Equal(t, "alpha_numeric", result.Code)
		require.Equal(t, "foo.bar@baz", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AlphaNumericRegex.String()}, result.Metadata)
	})
}
