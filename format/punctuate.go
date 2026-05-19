package format

import "github.com/studiolambda/golidate"

// Punctuate returns a formatter that appends a period when one is missing.
//
// Empty strings are returned unchanged, and messages already ending in a period
// are not modified.
func Punctuate() golidate.Formatter {
	return func(message string) string {
		if len(message) == 0 {
			return message
		}

		last := rune(message[len(message)-1])

		if last != '.' && last != '!' && last != '?' {
			return message + "."
		}

		return message
	}
}
