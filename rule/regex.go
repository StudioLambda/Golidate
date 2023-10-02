package rule

import (
	"fmt"
	"regexp"

	"github.com/studiolambda/golidate"
)

func Regex(expression *regexp.Regexp) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "regex").
			With("regex", expression.String())

		val := fmt.Sprintf("%+v", value)

		if !expression.MatchString(val) {
			return result.Fail()
		}

		return result.Pass()
	}
}
