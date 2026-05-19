package rule_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

// TestMapKeys verifies map key child validation behavior.
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

	t.Run("DeterministicOrder", func(t *testing.T) {
		value := map[string]int{
			"c3": 3,
			"a1": 1,
			"b2": 2,
		}

		result := rule.MapKeys[map[string]int](
			rule.Alpha(),
		)(value)
		results := result.Name("keys").Results("keys").Failed()

		require.Len(t, results, 4)
		require.Equal(t, "keys", results[0].Attribute)
		require.Equal(t, "keys.[a1]", results[1].Attribute)
		require.Equal(t, "keys.[b2]", results[2].Attribute)
		require.Equal(t, "keys.[c3]", results[3].Attribute)
	})
}
