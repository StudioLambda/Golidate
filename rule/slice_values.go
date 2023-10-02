package rule

import (
	"fmt"

	"github.com/studiolambda/golidate"
)

func SliceValues[S ~[]T, T any](rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "slice_values")

		if iterable, ok := value.(S); ok {
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
