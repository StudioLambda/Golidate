package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

func Min(min int64) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "min").
			With("min", min)

		val, ok := numberValue(value)

		if !ok || val < float64(min) {
			return result.Fail()
		}

		return result.Pass()
	}
}

func numberValue(value any) (float64, bool) {
	reflected := reflect.ValueOf(value)

	if !reflected.IsValid() {
		return 0, false
	}

	switch reflected.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(reflected.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(reflected.Uint()), true
	case reflect.Float32, reflect.Float64:
		return reflected.Float(), true
	default:
		return 0, false
	}
}
