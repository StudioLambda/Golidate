package golidate

// Rule validates a value and returns a result.
//
// A rule should return a Result with a stable code and an explicit pass or fail
// state. The rule package provides constructors for common checks, and callers
// can wrap any Rule with methods such as Code, Message, With, and When.
type Rule func(value any) Result

// Rules is a list of validation rules.
//
// The type is primarily useful when an API wants to accept or store a reusable
// group of rules before passing them to Pending.Rules or Pending.AppendRules.
type Rules []Rule

// Code overrides the result code returned by a rule.
//
// The wrapped rule still performs the validation and keeps its metadata,
// message, pass state, and children. Only Result.Code is replaced.
func (rule Rule) Code(code string) Rule {
	return func(value any) Result {
		result := rule(value)
		result.Code = code

		return result
	}
}

// Message overrides the result message returned by a rule.
//
// The message is the untranslated text returned by Results.Messages unless a
// later translation entry replaces it.
func (rule Rule) Message(message string) Rule {
	return func(value any) Result {
		result := rule(value)
		result.Message = message

		return result
	}
}

// Rename overrides both the result code and message returned by a rule.
//
// Rename is a shorthand for Code(code).Message(code). Translation dictionaries
// usually use the code, while untranslated output exposes the message.
func (rule Rule) Rename(code string) Rule {
	return rule.
		Code(code).
		Message(code)
}

// With adds one metadata value to the result returned by a rule.
//
// Existing metadata is preserved unless the same key already exists, in which
// case the supplied value replaces it.
func (rule Rule) With(key string, metadataValue any) Rule {
	return func(value any) Result {
		return rule(value).With(key, metadataValue)
	}
}

// Conditions adds applicability conditions to the result returned by a rule.
//
// Conditions do not stop the wrapped rule from executing. They affect
// Result.Passes so conditionally irrelevant failures can be treated as passing.
func (rule Rule) Conditions(conditions ...Condition) Rule {
	return func(value any) Result {
		return rule(value).Conditions(conditions...)
	}
}

// When adds a boolean applicability condition to the rule result.
//
// When is convenient when the applicability decision is already known before
// the rule runs. A false condition makes Result.Passes return true.
func (rule Rule) When(condition bool) Rule {
	return func(value any) Result {
		return rule(value).When(condition)
	}
}

// WithMetadata replaces metadata on the result returned by a rule.
//
// Replacement discards any metadata produced by the wrapped rule. Use
// WithMetadataMerged when the existing rule metadata should remain available to
// translation entries.
func (rule Rule) WithMetadata(metadata Metadata) Rule {
	return func(value any) Result {
		return rule(value).WithMetadata(metadata)
	}
}

// WithMetadataMerged merges metadata into the result returned by a rule.
//
// Keys in the supplied metadata replace matching keys already present on the
// result. A nil metadata map on the result is initialized before merging.
func (rule Rule) WithMetadataMerged(metadata Metadata) Rule {
	return func(value any) Result {
		return rule(value).WithMetadataMerged(metadata)
	}
}
