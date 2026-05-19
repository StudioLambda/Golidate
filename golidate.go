package golidate

import "context"

// Validator validates itself and returns validation results.
type Validator interface {
	Validate(ctx context.Context) Results
}

// Validate runs all pending validations and returns their results.
func Validate(ctx context.Context, pendings ...Pending) Results {
	results := Results{}

	for _, pending := range pendings {
		results = append(results, pending.Validate(ctx)...)
	}

	return results
}
