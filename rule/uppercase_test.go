package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestUppercase(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Uppercase()("FOO BAR BAZ")

		require.True(t, result.Passes())
		require.Equal(t, "uppercase", result.Code)
		require.Equal(t, "FOO BAR BAZ", result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Uppercase()("Foo Bar Baz")

		require.False(t, result.Passes())
		require.Equal(t, "uppercase", result.Code)
		require.Equal(t, "Foo Bar Baz", result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})
}
