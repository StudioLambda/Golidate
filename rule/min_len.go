package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func MinLen(min int) golidate.Rule {
	return func(value any) (result golidate.Result) {
		result = golidate.
			Uncertain(value, "min_len").
			With("min", min)

		ref := reflect.ValueOf(value)

		defer func() {
			if r := recover(); r != nil {
				result.Fail()
			}
		}()

		if ref.Len() < min {
			return result.Fail()
		}

		return result.Pass()
	}
}
