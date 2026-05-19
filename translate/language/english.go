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
// When Not wraps a compound rule (And or Or), De Morgan's law is applied: And
// becomes "or else" and Or becomes "and also" so negation distributes correctly.
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
//
// Compound operations (And, Or) are handled with De Morgan's law: an And
// operation under Not becomes "or else" joined, and an Or operation becomes
// "and also" joined. This produces logically correct negated conjunctions.
//
// When the resulting message does not contain a standalone "must" token, Invert
// prefixes the message with "must not" so the negation is always visible.
func Invert(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
	message := result.Message

	if operation, ok := result.Metadata["operation"].(golidate.Result); ok {
		translated := invertCompound(dictionary, operation)
		message = translated.Message
		result.Metadata["operation"] = golidate.Result(translated)
	}

	inverted := invertMust(message)

	if inverted == message && message != "" {
		inverted = "must not " + message
	}

	result.Message = inverted

	return result
}

// invertCompound applies De Morgan's law when negating compound rules.
//
// And operations are re-joined with " or else " and Or operations with
// " and also ". Non-compound operations are translated normally.
func invertCompound(dictionary golidate.Dictionary, operation golidate.Result) golidate.Result {
	switch operation.Code {
	case "and":
		return translate.SplitFromMetadata("operations", " or else ")(dictionary, operation)
	case "or":
		return translate.SplitFromMetadata("operations", " and also ")(dictionary, operation)
	default:
		return operation.Translate(dictionary)
	}
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
