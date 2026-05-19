package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

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
