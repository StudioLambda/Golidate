package rule

import "github.com/studiolambda/golidate"

func Type[T any]() golidate.Rule {
	return TypeOf(*new(T)).Rename("type")
}
