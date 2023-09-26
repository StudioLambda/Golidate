package language

import (
	"strings"

	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/translate"
)

// The English language for golidate. Please
// note that there are some special keywords
// that are used to split the message into multiple
// parts or inverse its meaning.
//
// These keywords are: `and also`, `or else`, `must`, `must not`
//
// Try finding alternative words if there is the need to write those
// keywords in the message without triggering the special behavior.`
var English = golidate.Dictionary{
	"accepted":       translate.Simple("the :attribute field must be accepted"),
	"alpha":          translate.Simple("the :attribute field must only contain letters"),
	"alpha_dash":     translate.Simple("the :attribute field must only contain letters, numbers, dashes, and underscores"),
	"alpha_extended": translate.Simple("the :attribute field must only contain letters, numbers, dashes, underscores and dots"),
	"alpha_numeric":  translate.Simple("the :attribute field must only contain letters and numbers"),
	"and":            translate.SplitFromMetadata("operations", " and also "),
	"ascii":          translate.Simple("the :attribute field must only contain single-byte alphanumeric characters and symbols"),
	"boolean":        translate.Simple("the :attribute field must be @values"),
	"convertible":    translate.Simple("the :attribute field must be convertible to @type"),
	"declined":       translate.Simple("the :attribute field must be declined"),
	"domain":         translate.Simple("the :attribute field must be a valid domain"),
	"email":          translate.Simple("the :attribute field must be a valid email address"),
	"equal":          translate.Simple("the :attribute field must be equal to @value"),
	"has_prefix":     translate.Simple("the :attribute field must start with @prefix"),
	"has_suffix":     translate.Simple("the :attribute field must end with @suffix"),
	"in":             translate.Simple("the :attribute field must be one of @values"),
	"len":            translate.Simple("the :attribute field must be exactly @len characters long"),
	"lowercase":      translate.Simple("the :attribute field must only contain lowercase letters"),
	"max":            translate.Simple("the :attribute field must be at most @max"),
	"max_len":        translate.Simple("the :attribute field must be at most @max characters long"),
	"min":            translate.Simple("the :attribute field must be at least @min"),
	"min_len":        translate.Simple("the :attribute field must be at least @min characters long"),
	"nil":            translate.Simple("the :attribute field must be nil"),
	"not":            Invert,
	"or":             translate.SplitFromMetadata("operations", " or else "),
	"regex":          translate.Simple("the :attribute field must match the regular expression @regex"),
	"slice_each":     translate.Simple("the :attribute field must be a slice"),
	"type":           translate.Simple("the :attribute field must be of type @type"),
	"type_of":        translate.Simple("the :attribute field must be of type @type"),
	"uppercase":      translate.Simple("the :attribute field must only contain uppercase letters"),
	"url":            translate.Simple("the :attribute field must be a valid URL"),
	"zero":           translate.Simple("the :attribute field must be a zero value"),
}

func Invert(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
	message := result.Message

	if operation, ok := result.Metadata["operation"].(golidate.Result); ok {
		translated := operation.Translate(dictionary)
		message = translated.Message
		result.Metadata["operation"] = golidate.Result(translated)
	}

	result.Message = strings.ReplaceAll(message, "must", "must not")
	result.Message = strings.ReplaceAll(result.Message, "not not", "")

	return result
}
