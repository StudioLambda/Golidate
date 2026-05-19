// Package rule provides common validation rules for golidate.
//
// Rules in this package return golidate.Result values with stable codes and
// metadata suitable for translation. Most rules fail when a value has an
// incompatible type rather than attempting broad coercion. Text-oriented rules
// format values with fmt.Sprintf before matching, while typed, numeric, length,
// and collection rules document their stricter behavior individually.
//
// Custom rules must have matching translation dictionary entries to produce
// user-facing messages. Without a dictionary entry for the result code,
// translated output surfaces the raw code string unchanged. Register entries in
// a golidate.Dictionary keyed by the code returned in Result.Code.
package rule
