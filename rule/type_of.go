package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

// TypeOf returns a rule that passes when value has the same dynamic type as t.
//
// Nil expected or actual types fail safely. The expected type string is stored
// in metadata as "type" when it can be determined.
func TypeOf(t any) golidate.Rule {
	return func(value any) golidate.Result {
		of := reflect.TypeOf(t)
		ref := reflect.TypeOf(value)
		var typ any
		if of != nil {
			typ = of.String()
		}

		result := golidate.
			Uncertain(value, "type_of").
			With("type", typ)

		if ref == nil || of == nil || ref != of {
			return result.Fail()
		}

		return result.Pass()
	}
}
