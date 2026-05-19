package rule

import "github.com/studiolambda/golidate"

// declinedValues stores the canonical string values accepted by Declined.
var declinedValues = []any{"false", "0", "off", "no"}

// DeclinedValues returns a defensive copy of Declined's allowed values.
var DeclinedValues = values(declinedValues)

// Declined returns a rule that passes for common negative string values.
//
// Declined uses strict membership, so the input must be a string equal to
// "false", "0", "off", or "no".
func Declined() golidate.Rule {
	return In(values(declinedValues)...).Rename("declined")
}
