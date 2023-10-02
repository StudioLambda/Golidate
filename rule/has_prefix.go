package rule

import (
	"fmt"
	"strings"

	"github.com/studiolambda/golidate"
)

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
