package language_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/translate"
	"github.com/studiolambda/golidate/translate/language"
)

// TestEnglishInvertDoesNotMutateWordsContainingMust verifies standalone negation.
func TestEnglishInvertDoesNotMutateWordsContainingMust(t *testing.T) {
	dictionary := golidate.Dictionary{
		"custom": translate.Simple("the :attribute field must be mustard"),
	}
	result := golidate.
		Uncertain("yellow", "not").
		With("operation", golidate.Uncertain("yellow", "custom").Fail()).
		Fail()

	translated := language.Invert(dictionary, result)

	require.Equal(t, "the attribute field must not be mustard", translated.Message)
}

// TestEnglishInvertRemovesExistingNegation verifies negation can be toggled off.
func TestEnglishInvertRemovesExistingNegation(t *testing.T) {
	dictionary := golidate.Dictionary{
		"custom": translate.Simple("the :attribute field must not be empty"),
	}
	result := golidate.
		Uncertain("value", "not").
		With("operation", golidate.Uncertain("value", "custom").Fail()).
		Fail()

	translated := language.Invert(dictionary, result)

	require.Equal(t, "the attribute field must be empty", translated.Message)
}

// TestEnglishInvertNotAndProducesOr verifies De Morgan on And.
func TestEnglishInvertNotAndProducesOr(t *testing.T) {
	operations := golidate.Results{
		golidate.Uncertain("x", "min").With("min", 1).Fail(),
		golidate.Uncertain("x", "max").With("max", 10).Fail(),
	}

	andResult := golidate.
		Uncertain("x", "and").
		With("operations", operations).
		Fail()

	notResult := golidate.
		Uncertain("x", "not").
		With("operation", andResult).
		Fail()

	translated := notResult.Translate(language.English)

	require.Contains(t, translated.Message, " or else ")
	require.NotContains(t, translated.Message, " and also ")
}

// TestEnglishInvertNotOrProducesAnd verifies De Morgan on Or.
func TestEnglishInvertNotOrProducesAnd(t *testing.T) {
	operations := golidate.Results{
		golidate.Uncertain("x", "min").With("min", 1).Fail(),
		golidate.Uncertain("x", "max").With("max", 10).Fail(),
	}

	orResult := golidate.
		Uncertain("x", "or").
		With("operations", operations).
		Fail()

	notResult := golidate.
		Uncertain("x", "not").
		With("operation", orResult).
		Fail()

	translated := notResult.Translate(language.English)

	require.Contains(t, translated.Message, " and also ")
	require.NotContains(t, translated.Message, " or else ")
}

// TestEnglishInvertFallbackWithoutMust verifies messages without must get prefixed.
func TestEnglishInvertFallbackWithoutMust(t *testing.T) {
	dictionary := golidate.Dictionary{
		"custom": translate.Simple("something custom"),
	}

	result := golidate.
		Uncertain("v", "not").
		With("operation", golidate.Uncertain("v", "custom").Fail()).
		Fail()

	translated := language.Invert(dictionary, result)

	require.Equal(t, "must not something custom", translated.Message)
}
