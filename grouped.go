package golidate

// Grouped stores results by attribute name.
//
// Results.Group creates this map using Result.Attribute as the key. Each value
// preserves the original result order for that attribute.
type Grouped map[string]Results

// Messages returns grouped messages after applying optional formatters.
//
// The returned map has the same attribute keys as the grouped results. Each
// message is taken from Result.Message and then passed through every formatter
// in the order supplied.
func (grouped Grouped) Messages(formatters ...Formatter) map[string][]string {
	messages := make(map[string][]string)

	for key, results := range grouped {
		messages[key] = results.Messages(formatters...)
	}

	return messages
}
