package rule

import (
	"strings"

	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func HasPrefix(prefix string) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "has_prefix").
			With("prefix", prefix)

		val, err := cast.ToStringE(value)

		if err != nil || !strings.HasPrefix(val, prefix) {
			return result.Fail()
		}

		return result.Pass()
	}
}
