package rule

import (
	"fmt"

	"github.com/studiolambda/golidate"
)

// SliceValues returns a rule that applies child rules to every slice element.
//
// The value must have the exact slice type S. Each child result is prefixed with
// its numeric index, producing attributes such as tags.0 after the parent result
// is expanded.
func SliceValues[S ~[]T, T any](rules ...golidate.Rule) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "slice_values")

		if iterable, ok := value.(S); ok {
			for _, rule := range rules {
				for i, current := range iterable {
					res := rule(current).Name(fmt.Sprintf("%d", i))
					result = result.WithPrefixedChild(res)
				}
			}

			return result.Pass()
		}

		return result.Fail()
	}
}
