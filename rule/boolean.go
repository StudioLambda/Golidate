package rule

import "github.com/studiolambda/golidate"

var booleanValues = []any{"true", "false", "1", "0", "on", "off", "yes", "no"}

var BooleanValues = values(booleanValues)

func Boolean() golidate.Rule {
	return In(values(booleanValues)...).Rename("boolean")
}
