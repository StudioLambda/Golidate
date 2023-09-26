package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestDomain(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Domain()("erik.cat")

		require.True(t, result.Passes())
		require.Equal(t, "domain", result.Code)
		require.Equal(t, "erik.cat", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.DomainRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Domain()("bad-domain")

		require.False(t, result.Passes())
		require.Equal(t, "domain", result.Code)
		require.Equal(t, "bad-domain", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.DomainRegex.String()}, result.Metadata)
	})
}
