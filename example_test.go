package golidate_test

import (
	"context"
	"fmt"

	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
	"github.com/studiolambda/golidate/translate/language"
)

// ExampleValidate_values demonstrates validating one value and translating failures.
func ExampleValidate_values() {
	results := golidate.Validate(context.Background()).Values(
		golidate.Value("erik42").Name("username").Rules(
			rule.Alpha(),
		),
	)

	for _, message := range results.Failed().Translate(language.English).Messages() {
		fmt.Println(message)
	}

	// Output:
	// the username field must only contain letters
}
