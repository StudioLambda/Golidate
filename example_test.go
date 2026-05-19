package golidate_test

import (
	"context"
	"fmt"

	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
	"github.com/studiolambda/golidate/translate/language"
)

func ExampleValidate() {
	results := golidate.Validate(
		context.Background(),
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
