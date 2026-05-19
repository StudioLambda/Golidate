// Package golidate validates Go values with small, composable rules.
//
// A validation starts with Value or Self, receives one or more Rule values, and
// produces Result values. Results retain the original value, a stable code, a
// display message, optional metadata, and any nested child results created by
// collection or composite rules.
//
// Values that implement Validator are validated recursively when passed through
// Value. Self disables that recursive Validator shortcut and lets callers apply
// rules to the value itself. Pointers to validators are supported, and map
// traversal is sorted by formatted key text so repeated validations produce
// deterministic result ordering.
package golidate
