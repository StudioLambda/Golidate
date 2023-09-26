package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestSliceEach(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		value := []int{10, 15, 20}
		result := rule.SliceEach[int](
			rule.Min(10),
			rule.Max(20),
		)(value)

		require.True(t, result.PassesAll())
		require.Equal(t, "slice_each", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})

	t.Run("Failure", func(t *testing.T) {
		value := []int{5, 15, 25}
		result := rule.SliceEach[int](
			rule.Min(10),
			rule.Max(20),
		)(value)

		require.True(t, result.Passes())
		require.False(t, result.PassesChilds())
		require.Equal(t, "slice_each", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})
}
