package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

// Convertible returns a rule that passes when value can be converted to T.
//
// Nil actual or target types fail safely. The rule uses reflect.Type.ConvertibleTo
// and stores the target type string in metadata as "type" when available.
func Convertible[T any]() golidate.Rule {
	return func(value any) golidate.Result {
		ref := reflect.TypeOf(value)
		of := reflect.TypeOf(*new(T))
		var typ any
		if of != nil {
			typ = of.String()
		}

		result := golidate.
			Uncertain(value, "convertible").
			With("type", typ)

		if ref == nil || of == nil || !ref.ConvertibleTo(of) {
			return result.Fail()
		}

		return result.Pass()
	}
}
