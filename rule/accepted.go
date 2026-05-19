package rule

import "github.com/studiolambda/golidate"

// acceptedValues stores the canonical string values accepted by Accepted.
var acceptedValues = []any{"true", "1", "on", "yes"}

// AcceptedValues returns a defensive copy of Accepted's allowed values.
var AcceptedValues = values(acceptedValues)

// Accepted returns a rule that passes for common affirmative string values.
//
// Accepted uses the same strict membership behavior as In, so the input must be
// a string equal to one of "true", "1", "on", or "yes".
func Accepted() golidate.Rule {
	return In(values(acceptedValues)...).Rename("accepted")
}
