package format

import (
	"strings"
	"unicode"

	"github.com/studiolambda/golidate"
)

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

func firstRune(message string) (rune, int) {
	reader := strings.NewReader(message)
	r, size, err := reader.ReadRune()

	if err != nil {
		return 0, 0
	}

	return r, size
}
