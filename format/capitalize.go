package format

import (
	"strings"
	"unicode"

	"github.com/studiolambda/golidate"
)

// Capitalize returns a formatter that uppercases the first rune of a message.
//
// Empty strings are returned unchanged. The formatter is Unicode-aware for the
// first rune and leaves the remainder of the message untouched.
func Capitalize() golidate.Formatter {
	return func(message string) string {
		if len(message) == 0 {
			return message
		}

		r, size := firstRune(message)
		uppercased := unicode.ToUpper(r)

		return string(uppercased) + message[size:]
	}
}

// firstRune returns the first UTF-8 rune and its byte size.
//
// Invalid or empty input returns zero values, which lets callers keep the
// original string unchanged when no first rune can be read.
func firstRune(message string) (rune, int) {
	reader := strings.NewReader(message)
	r, size, err := reader.ReadRune()

	if err != nil {
		return 0, 0
	}

	return r, size
}
