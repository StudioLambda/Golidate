package rule

import (
	"fmt"
	"strings"

	"github.com/studiolambda/golidate"
)

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
