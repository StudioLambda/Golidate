package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestMapValues(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		value := map[string]int{
			"one":   10,
			"two":   15,
			"three": 20,
		}

		result := rule.MapValues[map[string]int](
			rule.Min(10),
			rule.Max(20),
		)(value)

		require.True(t, result.PassesAll())
		require.Equal(t, "map_values", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})

	t.Run("Failure", func(t *testing.T) {
		value := map[string]int{
			"one":   5,
			"two":   15,
			"three": 25,
		}

		result := rule.MapValues[map[string]int](
			rule.Min(10),
			rule.Max(20),
		)(value)

		require.True(t, result.Passes())
		require.False(t, result.PassesChilds())
		require.Equal(t, "map_values", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})

	t.Run("BadType", func(t *testing.T) {
		value := "bad type"

		result := rule.MapValues[map[string]int](
			rule.Min(10),
			rule.Max(20),
		)(value)

		require.False(t, result.Passes())
		require.True(t, result.PassesChilds())
		require.Equal(t, "map_values", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})
}
