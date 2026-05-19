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
