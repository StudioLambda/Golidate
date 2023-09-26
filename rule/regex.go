package rule

import (
	"regexp"

	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func Regex(expression *regexp.Regexp) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "regex").
			With("regex", expression.String())

		val, err := cast.ToStringE(value)

		if err != nil || !expression.MatchString(val) {
			return result.Fail()
		}

		return result.Pass()
	}
}
