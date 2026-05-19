package format

import (
	"unicode"

	"github.com/studiolambda/golidate"
)

// Uncapitalize returns a formatter that lowercases the first rune of a message.
//
// Empty strings are returned unchanged. The formatter is Unicode-aware for the
// first rune and leaves the remainder of the message untouched.
func Uncapitalize() golidate.Formatter {
	return func(message string) string {
		if len(message) == 0 {
			return message
		}

		r, size := firstRune(message)
		lowercased := unicode.ToLower(r)

		return string(lowercased) + message[size:]
	}
}
