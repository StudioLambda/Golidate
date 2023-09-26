package rule

import (
	"unicode"

	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func Uppercase() golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "uppercase")

		val, err := cast.ToStringE(value)

		if err != nil {
			return result.Fail()
		}

		for _, r := range val {
			if !unicode.IsUpper(r) && unicode.IsLetter(r) {
				return result.Fail()
			}
		}

		return result.Pass()
	}
}
