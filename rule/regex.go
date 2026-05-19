package rule

import (
	"fmt"
	"regexp"

	"github.com/studiolambda/golidate"
)

// Regex returns a rule that passes when expression matches the formatted value.
//
// Values are converted with fmt.Sprintf("%+v", value) before matching. The
// regular expression string is stored in metadata as "regex" for translators.
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
