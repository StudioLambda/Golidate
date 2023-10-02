package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestMapKeys(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		value := map[string]int{
			"one":   10,
			"two":   15,
			"three": 20,
		}

		result := rule.MapKeys[map[string]int](
			rule.Alpha(),
		)(value)

		require.True(t, result.PassesAll())
		require.Equal(t, "map_keys", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})

	t.Run("Failure", func(t *testing.T) {
		value := map[string]int{
			"one123":   5,
			"two":      15,
			"three321": 25,
		}

		result := rule.MapKeys[map[string]int](
			rule.Alpha(),
		)(value)

		require.True(t, result.Passes())
		require.False(t, result.PassesChilds())
		require.Equal(t, "map_keys", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})

	t.Run("BadType", func(t *testing.T) {
		value := "bad type"

		result := rule.MapKeys[map[string]int](
			rule.Alpha(),
		)(value)

		require.False(t, result.Passes())
		require.True(t, result.PassesChilds())
		require.Equal(t, "map_keys", result.Code)
		require.Equal(t, value, result.Value)
		require.Equal(t, golidate.Metadata{}, result.Metadata)
	})
}
