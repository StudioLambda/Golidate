package rule

import (
	"reflect"

	"github.com/studiolambda/golidate"
)

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
