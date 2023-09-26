package rule_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestRegex(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		expression := regexp.MustCompile(`^[a-z]+$`)
		result := rule.Regex(expression)("hello")

		require.True(t, result.Passes())
		require.Equal(t, "regex", result.Code)
		require.Equal(t, "hello", result.Value)
		require.Equal(t, golidate.Metadata{"regex": expression.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		expression := regexp.MustCompile(`^[a-z]+$`)
		result := rule.Regex(expression)("hello123")

		require.False(t, result.Passes())
		require.Equal(t, "regex", result.Code)
		require.Equal(t, "hello123", result.Value)
		require.Equal(t, golidate.Metadata{"regex": expression.String()}, result.Metadata)
	})
}
