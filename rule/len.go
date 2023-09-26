package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func Len(exact int) golidate.Rule {
	return func(value any) (result golidate.Result) {
		result = golidate.
			Uncertain(value, "len").
			With("len", exact)

		ref := reflect.ValueOf(value)

		defer func() {
			if r := recover(); r != nil {
				result.Fail()
			}
		}()

		if ref.Len() != exact {
			return result.Fail()
		}

		return result.Pass()
	}
}
