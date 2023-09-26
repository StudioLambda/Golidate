package golidate

type Pending struct {
	value     any
	attribute string
	rules     []Rule
}

func Value(value any) Pending {
	return Pending{
		value:     value,
		attribute: "",
		rules:     make([]Rule, 0),
	}
}

func (pending Pending) Rules(rules ...Rule) Pending {
	pending.rules = rules

	return pending
}

func (pending Pending) AppendRules(rules ...Rule) Pending {
	pending.rules = append(pending.rules, rules...)

	return pending
}

func (pending Pending) Name(attribute string) Pending {
	pending.attribute = attribute

	return pending
}

func (pending Pending) Validate() Results {
	results := make(Results, 0, len(pending.rules))

	for _, rule := range pending.rules {
		result := rule(pending.value)

		if pending.attribute != "" {
			result = result.Name(pending.attribute)
		}

		expanded := result.Results(pending.attribute)

		results = append(results, expanded...)
	}

	if validatable, ok := pending.value.(Validater); ok {
		results = append(results, validatable.Validate().Prefixed(pending.attribute)...)
	}

	return results
}
