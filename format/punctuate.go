package format

import "github.com/studiolambda/golidate"

func Punctuate() golidate.Formatter {
	return func(message string) string {
		if len(message) == 0 {
			return message
		}

		if message[len(message)-1] != '.' {
			return message + "."
		}

		return message
	}
}
