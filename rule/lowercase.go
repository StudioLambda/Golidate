package rule

import (
	"fmt"
	"unicode"

	"github.com/studiolambda/golidate"
)

// Lowercase returns a rule that passes when all letters are lowercase.
//
// Values are formatted with fmt.Sprintf before inspection. Non-letter runes are
// ignored, so digits, punctuation, symbols, and spaces do not cause a failure.
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
