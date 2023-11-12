package rule

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func Min(min int64) golidate.Rule {
	return func(value any) golidate.Result {
		result := golidate.
			Uncertain(value, "min").
			With("min", min)

		i, ok := value.(int)

		fmt.Printf("%T %T %T\n", value, i, ok)

		val, err := cast.ToInt64E(value)

		if err != nil || val < min {
			fmt.Println("Failed")
			return result.Fail()
		}

		fmt.Println("Ok")

		return result.Pass()
	}
}
