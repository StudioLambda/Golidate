package rule

import (
	"fmt"

	"github.com/studiolambda/golidate"
)

func InTyped[T any](values ...T) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "in_typed").
			With("values", values)

		val := fmt.Sprintf("%+v", value)

		for _, v := range values {
			if val == fmt.Sprintf("%+v", v) {
				return result.Pass()
			}
		}

		return result.Fail()
	}
}
