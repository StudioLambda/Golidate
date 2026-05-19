package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

// TestEqual verifies DeepEqual-based comparison behavior.
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

	t.Run("SlicePass", func(t *testing.T) {
		value := []int{1, 2}
		other := []int{1, 2}

		result := rule.Equal(other)(value)

		require.True(t, result.Passes())
		require.Equal(t, "equal", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"other": other}, result.Metadata)
	})

	t.Run("SliceFail", func(t *testing.T) {
		value := []int{1, 2}
		other := []int{2, 1}

		result := rule.Equal(other)(value)

		require.False(t, result.Passes())
		require.Equal(t, "equal", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"other": other}, result.Metadata)
	})

	t.Run("MapPass", func(t *testing.T) {
		value := map[string]int{"foo": 1}
		other := map[string]int{"foo": 1}

		result := rule.Equal(other)(value)

		require.True(t, result.Passes())
		require.Equal(t, "equal", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"other": other}, result.Metadata)
	})

	t.Run("MapFail", func(t *testing.T) {
		value := map[string]int{"foo": 1}
		other := map[string]int{"foo": 2}

		result := rule.Equal(other)(value)

		require.False(t, result.Passes())
		require.Equal(t, "equal", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{"other": other}, result.Metadata)
	})
}
