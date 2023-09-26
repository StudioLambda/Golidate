package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestAscii(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Ascii()("example")

		require.True(t, result.Passes())
		require.Equal(t, "ascii", result.Code)
		require.Equal(t, "example", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AsciiRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Ascii()("∞")

		require.False(t, result.Passes())
		require.Equal(t, "ascii", result.Code)
		require.Equal(t, "∞", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.AsciiRegex.String()}, result.Metadata)
	})
}
