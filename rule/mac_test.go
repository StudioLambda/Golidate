package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestMac(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Mac()("00:00:00:00:00:00")

		require.True(t, result.PassesAll())
		require.Equal(t, "mac", result.Code)
		require.Equal(t, "00:00:00:00:00:00", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.MacRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Mac()("foo.bar")

		require.False(t, result.PassesAll())
		require.Equal(t, "mac", result.Code)
		require.Equal(t, "foo.bar", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.MacRegex.String()}, result.Metadata)
	})
}
