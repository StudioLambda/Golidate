package golidate

import "context"

type Validator interface {
	Validate(ctx context.Context) Results
}

func Validate(ctx context.Context, pendings ...Pending) Results {
	results := Results{}

	for _, pending := range pendings {
		results = append(results, pending.Validate(ctx)...)
	}

	return results
}
