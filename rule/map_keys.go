package rule

import (
	"fmt"
	"sort"

	"github.com/studiolambda/golidate"
)

type formattedMapKey[K comparable] struct {
	name string
	key  K
}

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
