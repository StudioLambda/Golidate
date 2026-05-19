package golidate

// Formatter transforms validation messages for display.
//
// Formatters are applied by Results.Messages and Grouped.Messages after any
// translation has already updated Result.Message. They should be pure string
// transformations, such as capitalization or punctuation.
type Formatter func(string) string
