package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func In(values ...any) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "in").
			With("values", values)

		valueType := reflect.TypeOf(value)
		for _, v := range values {
			if valueType == reflect.TypeOf(v) && reflect.DeepEqual(value, v) {
				return result.Pass()
			}
		}

		return result.Fail()
	}
}
