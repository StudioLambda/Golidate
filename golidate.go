package golidate

type Validator interface {
	Validate() Results
}

func Validate(pendings ...Pending) Results {
	results := Results{}

	for _, pending := range pendings {
		results = append(results, pending.Validate()...)
	}

	return results
}
