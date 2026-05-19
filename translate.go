package golidate

// Dictionary maps result codes to translation entries.
//
// Later dictionaries passed to Result.Translate or Results.Translate override
// earlier dictionaries for matching codes. Missing codes leave the result
// unchanged, which makes partial custom dictionaries safe to layer on top of a
// built-in language dictionary.
type Dictionary map[string]Entry

// Entry translates a result using a dictionary.
//
// Entries receive the merged dictionary so compound entries can translate
// nested operation metadata with the same language rules as the outer result.
type Entry func(dictionary Dictionary, result Result) Result
