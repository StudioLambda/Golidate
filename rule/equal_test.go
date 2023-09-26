package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestEqual(t *testing.T) {
	t.Run("Pass", func(t *testing.T) {
		result := rule.Equal("foo")("foo")

		require.True(t, result.Passes())
		require.Equal(t, "equal", result.Code)
		require.Equal(t, "foo", result.Value)
		require.Equal(t, golidate.Metadata{"other": "foo"}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Equal("foo")("bar")

		require.False(t, result.Passes())
		require.Equal(t, "equal", result.Code)
		require.Equal(t, "bar", result.Value)
		require.Equal(t, golidate.Metadata{"other": "foo"}, result.Metadata)
	})
}
