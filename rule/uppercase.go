package rule

import (
	"fmt"
	"unicode"

	"github.com/studiolambda/golidate"
)

func Uppercase() golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "uppercase")

		val := fmt.Sprintf("%+v", value)

		for _, r := range val {
			if !unicode.IsUpper(r) && unicode.IsLetter(r) {
				return result.Fail()
			}
		}

		return result.Pass()
	}
}
