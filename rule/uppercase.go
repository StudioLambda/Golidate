package rule

import (
	"fmt"
	"unicode"

	"github.com/studiolambda/golidate"
)

// Uppercase returns a rule that passes when all letters are uppercase.
//
// Values are formatted with fmt.Sprintf before inspection. Non-letter runes are
// ignored, so digits, punctuation, symbols, and spaces do not cause a failure.
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
