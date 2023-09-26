package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestLowercase(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Lowercase()("foo bar baz")

		require.True(t, result.Passes())
		require.Equal(t, "lowercase", result.Code)
		require.Equal(t, "foo bar baz", result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Lowercase()("Foo Bar Baz")

		require.False(t, result.Passes())
		require.Equal(t, "lowercase", result.Code)
		require.Equal(t, "Foo Bar Baz", result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})
}
