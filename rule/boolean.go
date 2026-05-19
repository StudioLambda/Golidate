package rule

import "github.com/studiolambda/golidate"

// booleanValues stores the canonical string values accepted by Boolean.
var booleanValues = []any{"true", "false", "1", "0", "on", "off", "yes", "no"}

// BooleanValues returns a defensive copy of Boolean's allowed values.
var BooleanValues = values(booleanValues)

// Boolean returns a rule that passes for common boolean-like string values.
//
// Boolean does not accept the Go bool type. It uses strict membership against
// the string values "true", "false", "1", "0", "on", "off", "yes", and "no".
func Boolean() golidate.Rule {
	return In(values(booleanValues)...).Rename("boolean")
}
