package golidate

type Rule func(value any) Result

type Rules []Rule

func (rule Rule) Code(message string) Rule {
	return func(value any) Result {
		result := rule(value)
		result.Code = message

		return result
	}
}

func (rule Rule) With(key string, val any) Rule {
	return func(value any) Result {
		return rule(value).With(key, val)
	}
}

func (rule Rule) Conditions(conditions ...Condition) Rule {
	return func(value any) Result {
		return rule(value).Conditions(conditions...)
	}
}

func (rule Rule) When(condition bool) Rule {
	return func(value any) Result {
		return rule(value).When(condition)
	}
}

func (rule Rule) WithMetadata(metadata Metadata) Rule {
	return func(value any) Result {
		return rule(value).WithMetadata(metadata)
	}
}

func (rule Rule) WithMetadataMerged(metadata Metadata) Rule {
	return func(value any) Result {
		return rule(value).WithMetadataMerged(metadata)
	}
}
