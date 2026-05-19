package golidate

// Dictionary maps result codes to translation entries.
type Dictionary map[string]Entry

// Entry translates a result using a dictionary.
type Entry func(dictionary Dictionary, result Result) Result
