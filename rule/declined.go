package rule

import "github.com/studiolambda/golidate"

var declinedValues = []any{"false", "0", "off", "no"}

var DeclinedValues = values(declinedValues)

func Declined() golidate.Rule {
	return In(values(declinedValues)...).Rename("declined")
}
