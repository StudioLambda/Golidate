package golidate

// Results is a list of validation results.
//
// Result order is meaningful and generally follows the order in which pending
// validations and rules were supplied. Map-derived nested results are sorted by
// formatted key text before they are appended.
type Results []Result

// PassesAll reports whether all results pass, including child results.
//
// An empty Results value passes because no failing validation is present.
func (results Results) PassesAll() bool {
	for _, result := range results {
		if !result.PassesAll() {
			return false
		}
	}

	return true
}

// PassesAny reports whether any result passes, including child results.
//
// An empty Results value returns false because there is no passing validation to
// report.
func (results Results) PassesAny() bool {
	for _, result := range results {
		if result.PassesAll() {
			return true
		}
	}

	return false
}

// Failed returns results that failed, including child result state.
//
// The returned slice preserves the receiver order and keeps each Result intact,
// including metadata and unexpanded child results.
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
//
// The returned slice preserves the receiver order.
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
//
// An empty prefix returns the original slice unchanged. Non-empty prefixes are
// joined to existing attributes with a dot.
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
//
// Messages does not translate. Call Translate first when message codes should
// be expanded into user-facing language.
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
//
// Dictionaries are merged once for the entire slice, so layered translations
// apply consistently and efficiently across all results.
func (results Results) Translate(dictionaries ...Dictionary) Results {
	res := make(Results, len(results))
	dictionary := mergeDictionaries(dictionaries...)

	for i, result := range results {
		res[i] = result.translate(dictionary)
	}

	return res
}

// Group groups results by attribute name.
//
// Grouping is useful for form or API responses that need messages keyed by
// field. Result order is preserved inside each attribute group.
func (results Results) Group() Grouped {
	group := make(Grouped)

	for _, result := range results {
		group[result.Attribute] = append(group[result.Attribute], result)
	}

	return group
}

// Has reports whether any result has the given attribute name.
//
// Has checks only the results currently present in the slice. Call Result.Results
// or Pending.Validate first when nested children need to be flattened.
func (results Results) Has(attribute string) bool {
	for _, result := range results {
		if result.Attribute == attribute {
			return true
		}
	}

	return false
}
