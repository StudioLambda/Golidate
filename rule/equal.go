package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

// Equal returns a rule that passes when value deeply equals other.
//
// The rule uses reflect.DeepEqual, so slices, maps, pointers, and structs follow
// Go's DeepEqual semantics rather than string formatting or coercion.
func Equal(other any) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "equal").
			With("other", other)

		if !reflect.DeepEqual(value, other) {
			return result.Fail()
		}

		return result.Pass()
	}
}
