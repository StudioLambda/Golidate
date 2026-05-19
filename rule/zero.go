package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

// Zero returns a rule that passes when value is the zero value of its type.
//
// Invalid nil interfaces fail because no concrete type exists from which to
// derive a zero value. The expected zero value is stored in metadata as "zero".
func Zero() golidate.Rule {
	return func(value any) golidate.Result {
		val := reflect.ValueOf(value)
		if !val.IsValid() {
			return golidate.Uncertain(value, "zero").With("zero", nil).Fail()
		}

		zero := reflect.Zero(val.Type()).Interface()
		result := golidate.Uncertain(value, "zero").With("zero", zero)

		if !val.IsZero() {
			return result.Fail()
		}

		return result.Pass()
	}
}
