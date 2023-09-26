package rule

import (
	"fmt"

	"github.com/studiolambda/golidate"
)

func SliceEach[T any](rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "slice_each")

		if iterable, ok := value.([]T); ok {
			for _, rule := range rules {
				for i, current := range iterable {
					res := rule(current).Name(fmt.Sprintf("%d", i))
					result = result.WithChild(res)
				}
			}

			return result.Pass()
		}

		return result.Fail()
	}
}
