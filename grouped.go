package golidate

type Grouped map[string]Results

func (grouped Grouped) Messages(formatters ...Formatter) map[string][]string {
	messages := make(map[string][]string)

	for key, results := range grouped {
		messages[key] = results.Messages(formatters...)
	}

	return messages
}
