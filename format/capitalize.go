package format

import (
	"unicode"

	"github.com/studiolambda/golidate"
)

func Capitalize() golidate.Formatter {
	return func(message string) string {
		if len(message) == 0 {
			return message
		}

		uppercased := unicode.ToUpper(rune(message[0]))

		return string(uppercased) + message[1:]
	}
}
