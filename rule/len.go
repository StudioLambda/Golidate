package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

// Len returns a rule that passes when a value has exactly exact length.
//
// Arrays, channels, maps, slices, and strings are supported. Nil and unsupported
// values fail without panic recovery.
func Len(exact int) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "len").
			With("len", exact)

		length, ok := lengthOf(value)
		if !ok || length != exact {
			return result.Fail()
		}

		return result.Pass()
	}
}

// lengthOf returns the length for reflection kinds that support Len.
//
// Nil interfaces and unsupported kinds return false so length rules fail safely
// instead of relying on panic recovery.
func lengthOf(value any) (int, bool) {
	if value == nil {
		return 0, false
	}

	ref := reflect.ValueOf(value)
	switch ref.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return ref.Len(), true
	default:
		return 0, false
	}
}
