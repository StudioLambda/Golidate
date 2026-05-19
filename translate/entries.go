package translate

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cast"
	"github.com/studiolambda/golidate"
)

// Plain returns an entry that replaces the result message with fixed text.
//
// Plain ignores the dictionary and result metadata. It is useful for simple
// codes whose message does not need placeholders.
func Plain(message string) golidate.Entry {
	return func(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
		result.Message = message

		return result
	}
}

// SplitFromMetadata returns an entry that joins translated nested messages.
//
// The metadata value at key must be golidate.Results. Each nested result is
// translated with the same merged dictionary, stored back into metadata, and its
// message is joined with separator. Other metadata shapes leave the result
// unchanged.
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

// Simple returns an entry that expands common placeholders in a message.
//
// Supported colon placeholders are :attribute, :code, :message, and :value.
// Metadata placeholders use @key. Metadata keys are processed longest-first and
// then alphabetically so overlapping placeholders such as @min and @minimum are
// replaced deterministically.
func Simple(message string) golidate.Entry {
	return func(dictionary golidate.Dictionary, result golidate.Result) golidate.Result {
		message := strings.Clone(message)

		attribute := result.Attribute
		if attribute == "" {
			attribute = "attribute"
		}

		message = strings.ReplaceAll(message, ":attribute", attribute)
		message = strings.ReplaceAll(message, ":code", result.Code)
		message = strings.ReplaceAll(message, ":message", result.Message)

		value := fmt.Sprintf("%+v", result.Value)
		message = strings.ReplaceAll(message, ":value", value)

		for _, key := range metadataKeys(result.Metadata) {
			val := result.Metadata[key]

			if value, err := cast.ToStringSliceE(val); err == nil {
				message = strings.ReplaceAll(message, "@"+key, strings.Join(value, ", "))
				continue
			}

			value := fmt.Sprintf("%+v", val)
			message = strings.ReplaceAll(message, "@"+key, value)
		}

		result.Message = message

		return result
	}
}

// metadataKeys returns metadata keys in deterministic replacement order.
//
// Longer keys are returned before shorter keys so one placeholder cannot consume
// the prefix of another. Equal-length keys are sorted alphabetically for stable
// translation output.
func metadataKeys(metadata golidate.Metadata) []string {
	keys := make([]string, 0, len(metadata))

	for key := range metadata {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool {
		if len(keys[i]) == len(keys[j]) {
			return keys[i] < keys[j]
		}

		return len(keys[i]) > len(keys[j])
	})

	return keys
}
