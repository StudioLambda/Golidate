package rule

import (
	"github.com/studiolambda/golidate"
)

func InTyped[T comparable](values ...T) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "in_typed").
			With("values", values)

		val, ok := value.(T)
		if !ok {
			return result.Fail()
		}

		for _, v := range values {
			if val == v {
				return result.Pass()
			}
		}

		return result.Fail()
	}
}
