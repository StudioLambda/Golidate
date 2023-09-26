package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func MaxLen(max int) golidate.Rule {
	return func(value any) (result golidate.Result) {
		result = golidate.
			Uncertain(value, "max_len").
			With("max", max)

		ref := reflect.ValueOf(value)

		defer func() {
			if r := recover(); r != nil {
				result.Fail()
			}
		}()

		if ref.Len() > max {
			return result.Fail()
		}

		return result.Pass()
	}
}
