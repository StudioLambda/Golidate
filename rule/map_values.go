package rule

import (
	"fmt"

	"github.com/studiolambda/golidate"
)

func MapValues[M map[K]V, K comparable, V any](rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "map_values")

		if iterable, ok := value.(M); ok {
			for _, rule := range rules {
				for key, current := range iterable {
					res := rule(current).Name(fmt.Sprintf("%+v", key))
					result = result.WithPrefixedChild(res)
				}
			}

			return result.Pass()
		}

		return result.Fail()
	}
}
