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

		lowercased := unicode.ToLower(rune(message[0]))

		return string(lowercased) + message[1:]
	}
}
