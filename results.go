package golidate

// Results is a list of validation results.
type Results []Result

// PassesAll reports whether all results pass, including child results.
func (results Results) PassesAll() bool {
	for _, result := range results {
		if !result.PassesAll() {
			return false
		}
	}

	return true
}

// PassesAny reports whether any result passes, including child results.
func (results Results) PassesAny() bool {
	for _, result := range results {
		if result.PassesAll() {
			return true
		}
	}

	return false
}

// Failed returns results that failed, including child result state.
func (results Results) Failed() Results {
	failed := make(Results, 0, len(results))

	for _, result := range results {
		if !result.PassesAll() {
			failed = append(failed, result)
		}
	}

	return failed
}

// Passed returns results that passed, including child result state.
func (results Results) Passed() Results {
	passed := make(Results, 0, len(results))

	for _, result := range results {
		if result.PassesAll() {
			passed = append(passed, result)
		}
	}

	return passed
}

// Prefixed returns results with attribute names prefixed.
func (results Results) Prefixed(prefix string) Results {
	if prefix == "" {
		return results
	}

	prefixed := make(Results, len(results))

	for i, result := range results {
		prefixed[i] = result.Name(prefix + "." + result.Attribute)
	}

	return prefixed
}

// Messages returns result messages after applying optional formatters.
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

// Translate returns results translated by the provided dictionaries.
func (results Results) Translate(dictionaries ...Dictionary) Results {
	res := make(Results, len(results))
	dictionary := mergeDictionaries(dictionaries...)

	for i, result := range results {
		res[i] = result.translate(dictionary)
	}

	return res
}

// Group groups results by attribute name.
func (results Results) Group() Grouped {
	group := make(Grouped)

	for _, result := range results {
		group[result.Attribute] = append(group[result.Attribute], result)
	}

	return group
}

// Has reports whether any result has the given attribute name.
func (results Results) Has(attribute string) bool {
	for _, result := range results {
		if result.Attribute == attribute {
			return true
		}
	}

	return false
}
