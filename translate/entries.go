package translate

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

func Plain(message string) golidate.Entry {
	return func(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
		result.Message = message

		return result
	}
}

func SplitFromMetadata(key string, separator string) golidate.Entry {
	return func(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
		if results, ok := result.Metadata[key].(golidate.Results); ok {
			translations := make([]string, len(results))
			operations := make(golidate.Results, len(results))

			for i, operation := range results.Translate(dictionary) {
				operations[i] = operation
				translations[i] = operation.Message
			}

			result.Metadata[key] = operations
			result.Message = strings.Join(translations, separator)
		}

		return result
	}
}

func Simple(message string) golidate.Entry {
	return func(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
		message := strings.Clone(message)
		message = strings.ReplaceAll(message, ":attribute", result.Attribute)
		message = strings.ReplaceAll(message, ":code", result.Code)
		message = strings.ReplaceAll(message, ":message", result.Message)

		value := fmt.Sprintf("%+v", result.Value)
		message = strings.ReplaceAll(message, ":value", value)

		for key, val := range result.Metadata {
			if value, err := cast.ToStringSliceE(val); err == nil {
				message = strings.ReplaceAll(message, "@"+key, strings.Join(value, ", "))
				continue
			}

			if value, err := cast.ToStringE(val); err == nil {
				message = strings.ReplaceAll(message, "@"+key, value)
				continue
			}
		}

		result.Message = message

		return result
	}
}
