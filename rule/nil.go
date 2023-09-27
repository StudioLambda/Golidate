package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func Nil() golidate.Rule {
	return func(value any) (result golidate.Result) {
		result = golidate.Uncertain(value, "nil")
		reflected := reflect.ValueOf(value)

		defer func() {
			if r := recover(); r != nil {
				result.Fail()
			}
		}()

		if value != nil && !reflected.IsNil() {
			return result.Fail()
		}

		return result.Pass()
	}
}
