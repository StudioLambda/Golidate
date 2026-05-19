package format

import (
	"unicode"

	"github.com/studiolambda/golidate"
)

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
