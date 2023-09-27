package rule

import "github.com/studiolambda/golidate"

var DeclinedValues = []any{"false", "0", "off", "no"}

func Declined() golidate.Rule {
	return In(DeclinedValues...).Rename("declined")
}
