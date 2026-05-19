package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

// In returns a rule that passes when value strictly matches one allowed value.
//
// The comparison requires both the same dynamic type and reflect.DeepEqual. This
// means int(1) does not match int64(1), and string "1" does not match numeric
// 1. The full allowed list is stored in metadata as "values".
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
