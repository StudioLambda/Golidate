package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func TypeOf(t any) golidate.Rule {
	return func(value any) golidate.Result {
		of := reflect.TypeOf(t)
		ref := reflect.TypeOf(value)
		result := golidate.
			Uncertain(value, "type_of").
			With("type", of.String())

		if ref != of {
			return result.Fail()
		}

		return result.Pass()
	}
}
