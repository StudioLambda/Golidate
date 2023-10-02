package rule

import (
	"fmt"
	"unicode"

	"github.com/studiolambda/golidate"
)

func Lowercase() golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "lowercase")

		val := fmt.Sprintf("%+v", value)

		for _, r := range val {
			if !unicode.IsLower(r) && unicode.IsLetter(r) {
				return result.Fail()
			}
		}

		return result.Pass()
	}
}
