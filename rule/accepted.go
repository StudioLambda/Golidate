package rule

import "github.com/studiolambda/golidate"

var acceptedValues = []any{"true", "1", "on", "yes"}

var AcceptedValues = values(acceptedValues)

func Accepted() golidate.Rule {
	return In(values(acceptedValues)...).Rename("accepted")
}
