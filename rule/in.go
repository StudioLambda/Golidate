package rule

import (
	"github.com/studiolambda/golidate"
)

func In(values ...any) golidate.Rule {
	return InTyped(values...).Code("in")
}
