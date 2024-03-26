package rule

import (
	"fmt"

	"github.com/studiolambda/golidate"
)

func MapKeys[M ~map[K]V, K comparable, V any](rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "map_keys")

		if iterable, ok := value.(M); ok {
			for _, rule := range rules {
				for key := range iterable {
					res := rule(key).Name(fmt.Sprintf("[%+v]", key))
					result = result.WithPrefixedChild(res)
				}
			}

			return result.Pass()
		}

		return result.Fail()
	}
}
