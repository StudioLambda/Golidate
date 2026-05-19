package rule

import (
	"fmt"
	"sort"

	"github.com/studiolambda/golidate"
)

// formattedMapKey stores a map key with the text used for stable ordering.
type formattedMapKey[K comparable] struct {
	// name is the formatted key text used in attributes and sorting.
	name string
	// key is the original typed map key used to read from the map.
	key K
}

// MapKeys returns a rule that applies child rules to every key in a map.
//
// The value must have the exact map type M. Keys are formatted and sorted before
// validation so child results have deterministic order. Child attributes are
// bracketed key names such as settings.[enabled].
func MapKeys[M ~map[K]V, K comparable, V any](rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "map_keys")

		if iterable, ok := value.(M); ok {
			for _, rule := range rules {
				for _, key := range sortedFormattedMapKeys(iterable) {
					res := rule(key.key).Name("[" + key.name + "]")
					result = result.WithPrefixedChild(res)
				}
			}

			return result.Pass()
		}

		return result.Fail()
	}
}

// sortedFormattedMapKeys returns map keys in deterministic formatted order.
//
// Sorting by formatted names keeps map key validation stable despite Go's
// randomized map iteration order.
func sortedFormattedMapKeys[M ~map[K]V, K comparable, V any](iterable M) []formattedMapKey[K] {
	keys := make([]formattedMapKey[K], 0, len(iterable))

	for key := range iterable {
		keys = append(keys, formattedMapKey[K]{
			name: fmt.Sprintf("%+v", key),
			key:  key,
		})
	}

	sort.Slice(keys, func(i int, j int) bool {
		return keys[i].name < keys[j].name
	})

	return keys
}
