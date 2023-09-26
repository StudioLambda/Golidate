package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func Convertible[T any]() golidate.Rule {
	return func(value any) golidate.Result {
		ref := reflect.TypeOf(value)
		of := reflect.TypeOf(*new(T))
		result := golidate.
			Uncertain(value, "convertible").
			With("type", of.String())

		if !ref.ConvertibleTo(of) {
			return result.Fail()
		}

		return result.Pass()
	}
}
