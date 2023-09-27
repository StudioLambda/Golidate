package rule

import "github.com/studiolambda/golidate"

var BooleanValues = []any{"true", "false", "1", "0", "on", "off", "yes", "no"}

func Boolean() golidate.Rule {
	return In(BooleanValues...).Rename("boolean")
}
