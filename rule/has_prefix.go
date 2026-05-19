package rule

import (
	"fmt"
	"strings"

	"github.com/studiolambda/golidate"
)

// HasPrefix returns a rule that passes when the formatted value starts with prefix.
//
// Values are converted with fmt.Sprintf("%+v", value), so this rule is useful
// for display-oriented checks rather than strict string-only validation.
func HasPrefix(prefix string) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "has_prefix").
			With("prefix", prefix)

		val := fmt.Sprintf("%+v", value)

		if !strings.HasPrefix(val, prefix) {
			return result.Fail()
		}

		return result.Pass()
	}
}
