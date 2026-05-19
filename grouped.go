package golidate

// Grouped stores results by attribute name.
type Grouped map[string]Results

// Messages returns grouped messages after applying optional formatters.
func (grouped Grouped) Messages(formatters ...Formatter) map[string][]string {
	messages := make(map[string][]string)

	for key, results := range grouped {
		messages[key] = results.Messages(formatters...)
	}

	return messages
}
