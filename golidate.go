package golidate

import "context"

// Validator validates itself and returns validation results.
//
// Value detects Validator implementations on the value itself and on supported
// collection elements. Implementations should use the supplied context for any
// request-scoped values needed while building validations.
type Validator interface {
	// Validate returns every validation result for the receiver.
	//
	// Implementations commonly call golidate.Validate with child Pending values.
	// Returned attributes are prefixed by parent Value traversal when the
	// validator is nested inside a slice, array, map, or named field.
	Validate(ctx context.Context) Results
}

// Validate runs all pending validations and returns their results.
//
// Pendings are evaluated in the order provided. The returned slice preserves
// that order, except for nested map validation where map keys are sorted by
// their formatted names to avoid Go's randomized map iteration order.
func Validate(ctx context.Context, pendings ...Pending) Results {
	results := Results{}

	for _, pending := range pendings {
		results = append(results, pending.Validate(ctx)...)
	}

	return results
}
