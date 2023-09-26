package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestUrl(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result := rule.Url()("https://erik.cat/foo/bar#baz?name=john&last_name=doe")

		require.True(t, result.Passes())
		require.Equal(t, "url", result.Code)
		require.Equal(t, "https://erik.cat/foo/bar#baz?name=john&last_name=doe", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.UrlRegex.String()}, result.Metadata)
	})

	t.Run("Fail", func(t *testing.T) {
		result := rule.Url()("foo")

		require.False(t, result.Passes())
		require.Equal(t, "url", result.Code)
		require.Equal(t, "foo", result.Value)
		require.Equal(t, golidate.Metadata{"regex": rule.UrlRegex.String()}, result.Metadata)
	})
}
