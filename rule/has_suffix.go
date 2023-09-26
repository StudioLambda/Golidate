package rule

import (
	"strings"

	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func HasSuffix(suffix string) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "has_suffix").
			With("suffix", suffix)

		val, err := cast.ToStringE(value)

		if err != nil || !strings.HasSuffix(val, suffix) {
			return result.Fail()
		}

		return result.Pass()
	}
}
