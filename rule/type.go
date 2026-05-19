package rule

import "github.com/studiolambda/golidate"

// Type returns a rule that passes when value has exactly type T.
//
// The check delegates to TypeOf and then renames the code to "type" for
// translation compatibility. It does not accept values merely convertible to T.
func Type[T any]() golidate.Rule {
	return TypeOf(*new(T)).Rename("type")
}
