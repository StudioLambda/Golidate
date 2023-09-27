package rule

import "github.com/studiolambda/golidate"

var AcceptedValues = []any{"true", "1", "on", "yes"}

func Accepted() golidate.Rule {
	return In(AcceptedValues...).Rename("accepted")
}
