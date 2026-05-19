package golidate

// Rule validates a value and returns a result.
type Rule func(value any) Result

// Rules is a list of validation rules.
type Rules []Rule

// Code overrides the result code returned by a rule.
func (rule Rule) Code(code string) Rule {
	return func(value any) Result {
		result := rule(value)
		result.Code = code

		return result
	}
}

// Message overrides the result message returned by a rule.
func (rule Rule) Message(message string) Rule {
	return func(value any) Result {
		result := rule(value)
		result.Message = message

		return result
	}
}

// Rename overrides both the result code and message returned by a rule.
func (rule Rule) Rename(code string) Rule {
	return rule.
		Code(code).
		Message(code)
}

// With adds metadata to the result returned by a rule.
func (rule Rule) With(key string, val any) Rule {
	return func(value any) Result {
		return rule(value).With(key, val)
	}
}

// Conditions adds applicability conditions to the result returned by a rule.
func (rule Rule) Conditions(conditions ...Condition) Rule {
	return func(value any) Result {
		return rule(value).Conditions(conditions...)
	}
}

// When adds a boolean applicability condition to the rule result.
func (rule Rule) When(condition bool) Rule {
	return func(value any) Result {
		return rule(value).When(condition)
	}
}

// WithMetadata replaces metadata on the result returned by a rule.
func (rule Rule) WithMetadata(metadata Metadata) Rule {
	return func(value any) Result {
		return rule(value).WithMetadata(metadata)
	}
}

// WithMetadataMerged merges metadata into the result returned by a rule.
func (rule Rule) WithMetadataMerged(metadata Metadata) Rule {
	return func(value any) Result {
		return rule(value).WithMetadataMerged(metadata)
	}
}
