package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

// Nil returns a rule that passes when value is nil or a nil-capable nil value.
//
// Nil interfaces pass. Nil-capable values such as pointers, slices, maps,
// channels, and functions pass when their reflected value is nil. Non-nil values
// fail; unsupported IsNil calls are guarded by recovery for safety.
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
