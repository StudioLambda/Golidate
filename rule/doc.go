// Package rule provides common validation rules for golidate.
//
// Rules in this package return golidate.Result values with stable codes and
// metadata suitable for translation. Most rules fail when a value has an
// incompatible type rather than attempting broad coercion. Text-oriented rules
// format values with fmt.Sprintf before matching, while typed, numeric, length,
// and collection rules document their stricter behavior individually.
package rule
