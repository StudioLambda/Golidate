package language

import (
	"regexp"
	"strings"

	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/translate"
)

var (
	// mustPattern locates standalone "must" tokens for English negation.
	mustPattern = regexp.MustCompile(`\bmust\b`)
	// negatedMustPattern detects an existing "not" immediately after "must".
	negatedMustPattern = regexp.MustCompile(`^\s+not\b`)
)

// English is the built-in English translation dictionary for golidate.
//
// Compound entries use special joining phrases such as "and also" and "or else".
// The not entry inverts standalone "must" and "must not" phrases without
// changing words that merely contain those letters.
//
// These keywords are: `and also`, `or else`, `must`, `must not`
//
// Try alternate wording if a custom message needs those exact phrases without
// triggering compound or negation behavior.
var English = golidate.Dictionary{
	"accepted":       translate.Simple("the :attribute field must be one of @values"),
	"alpha":          translate.Simple("the :attribute field must only contain letters"),
	"alpha_dash":     translate.Simple("the :attribute field must only contain letters, numbers, dashes, and underscores"),
	"alpha_extended": translate.Simple("the :attribute field must only contain letters, numbers, dashes, underscores and dots"),
	"alpha_numeric":  translate.Simple("the :attribute field must only contain letters and numbers"),
	"and":            translate.SplitFromMetadata("operations", " and also "),
	"ascii":          translate.Simple("the :attribute field must only contain single-byte alphanumeric characters and symbols"),
	"boolean":        translate.Simple("the :attribute field must be one of @values"),
	"convertible":    translate.Simple("the :attribute field must be convertible to @type"),
	"declined":       translate.Simple("the :attribute field must be declined"),
	"domain":         translate.Simple("the :attribute field must be a valid domain"),
	"email":          translate.Simple("the :attribute field must be a valid email address"),
	"equal":          translate.Simple("the :attribute field must be equal to @value"),
	"has_prefix":     translate.Simple("the :attribute field must start with @prefix"),
	"has_suffix":     translate.Simple("the :attribute field must end with @suffix"),
	"in":             translate.Simple("the :attribute field must be one of @values"),
	"in_typed":       translate.Simple("the :attribute field must be one of @values"),
	"len":            translate.Simple("the :attribute field must be exactly @len characters long"),
	"lowercase":      translate.Simple("the :attribute field must only contain lowercase letters"),
	"mac":            translate.Simple("the :attribute field must be a valid MAC address"),
	"map_keys":       translate.Simple("the :attribute field must be a valid map"),
	"map_values":     translate.Simple("the :attribute field must be a valid map"),
	"max":            translate.Simple("the :attribute field must be at most @max"),
	"max_len":        translate.Simple("the :attribute field must be at most @max characters long"),
	"min":            translate.Simple("the :attribute field must be at least @min"),
	"min_len":        translate.Simple("the :attribute field must be at least @min characters long"),
	"nil":            translate.Simple("the :attribute field must be nil"),
	"not":            Invert,
	"optional":       translate.SplitFromMetadata("operations", " or else "),
	"or":             translate.SplitFromMetadata("operations", " or else "),
	"regex":          translate.Simple("the :attribute field must match the regular expression @regex"),
	"required":       translate.SplitFromMetadata("operations", " and also "),
	"slice_values":   translate.Simple("the :attribute field must be a valid slice"),
	"type":           translate.Simple("the :attribute field must be of type @type"),
	"type_of":        translate.Simple("the :attribute field must be of type @type"),
	"uppercase":      translate.Simple("the :attribute field must only contain uppercase letters"),
	"url":            translate.Simple("the :attribute field must be a valid URL"),
	"zero":           translate.Simple("the :attribute field must be a zero value"),
}

// Invert translates a nested operation and flips its English "must" wording.
//
// The nested operation is expected in result metadata under "operation". If it
// is present, the translated operation is written back into metadata so callers
// can inspect the same message that was used for inversion.
func Invert(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
	message := result.Message

	if operation, ok := result.Metadata["operation"].(golidate.Result); ok {
		translated := operation.Translate(dictionary)
		message = translated.Message
		result.Metadata["operation"] = golidate.Result(translated)
	}

	result.Message = invertMust(message)

	return result
}

// invertMust toggles standalone English "must" and "must not" phrases.
//
// Words such as "mustard" are intentionally ignored. Existing negation is
// removed, while a positive "must" receives "not".
func invertMust(message string) string {
	matches := mustPattern.FindAllStringIndex(message, -1)

	if len(matches) == 0 {
		return message
	}

	builder := strings.Builder{}
	last := 0

	for _, match := range matches {
		builder.WriteString(message[last:match[0]])

		if negated := negatedMustPattern.FindStringIndex(message[match[1]:]); negated != nil {
			builder.WriteString("must")
			last = match[1] + negated[1]

			continue
		}

		builder.WriteString("must not")
		last = match[1]
	}

	builder.WriteString(message[last:])

	return builder.String()
}
