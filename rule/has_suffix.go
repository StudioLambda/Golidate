package rule

import (
	"fmt"
	"strings"

	"github.com/studiolambda/golidate"
)

// HasSuffix returns a rule that passes when the formatted value ends with suffix.
//
// Values are converted with fmt.Sprintf("%+v", value), so this rule is useful
// for display-oriented checks rather than strict string-only validation.
func HasSuffix(suffix string) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "has_suffix").
			With("suffix", suffix)

		val := fmt.Sprintf("%+v", value)

		if !strings.HasSuffix(val, suffix) {
			return result.Fail()
		}

		return result.Pass()
	}
}
