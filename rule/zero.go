package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func Zero() golidate.Rule {
	return func(value any) golidate.Result {
		val := reflect.ValueOf(value)
		zero := reflect.Zero(val.Type()).Interface()
		result := golidate.Uncertain(value, "zero").With("zero", zero)

		if !val.IsZero() {
			return result.Fail()
		}

		return result.Pass()
	}
}
