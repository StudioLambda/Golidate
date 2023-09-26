package golidate

type Results []Result

func (results Results) PassesAll() bool {
	for _, result := range results {
		if !result.Passes() {
			return false
		}
	}

	return true
}

func (results Results) PassesAny() bool {
	for _, result := range results {
		if result.Passes() {
			return true
		}
	}

	return false
}

func (results Results) Failed() Results {
	failed := make(Results, 0, len(results))

	for _, result := range results {
		if !result.Passes() {
			failed = append(failed, result)
		}
	}

	return failed
}

func (results Results) Passed() Results {
	passed := make(Results, 0, len(results))

	for _, result := range results {
		if result.Passes() {
			passed = append(passed, result)
		}
	}

	return passed
}

func (results Results) Prefixed(prefix string) Results {
	prefixed := make(Results, len(results))

	for i, result := range results {
		prefixed[i] = result.Name(prefix + "." + result.Attribute)
	}

	return prefixed
}

func (results Results) Messages(formatters ...Formatter) []string {
	messages := make([]string, len(results))

	for i, result := range results {
		messages[i] = result.Message

		for _, formatter := range formatters {
			messages[i] = formatter(messages[i])
		}
	}

	return messages
}

func (results Results) Translate(dictionaries ...Dictionary) Results {
	res := make(Results, len(results))

	for i, result := range results {
		res[i] = result.Translate(dictionaries...)
	}

	return res
}

func (results Results) Group() Grouped {
	group := make(Grouped)

	for _, result := range results {
		group[result.Attribute] = append(group[result.Attribute], result)
	}

	return group
}
