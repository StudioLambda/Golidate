package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestEmail(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Email()("foo@bar.baz")

		require.True(t, result.Passes())
		require.Equal(t, "email", result.Code)
		require.Equal(t, "foo@bar.baz", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.EmailRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Email()("foo.bar")

		require.False(t, result.Passes())
		require.Equal(t, "email", result.Code)
		require.Equal(t, "foo.bar", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.EmailRegex.String()}, result.Metadata)
	})
}
