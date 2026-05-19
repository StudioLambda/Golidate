package rule

// values returns a copy of a package-level allowed-value slice.
//
// The copy prevents callers from mutating the backing array used by built-in
// rules such as Accepted, Boolean, and Declined.
func values(source []any) []any {
	return append([]any(nil), source...)
}
