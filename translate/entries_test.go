package translate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/translate"
	"github.com/studiolambda/golidate/translate/language"
)

// TestSimple verifies placeholder replacement for attributes, codes, values, and metadata.
func TestSimple(t *testing.T) {
	result := golidate.
		Uncertain(10, "simple").
		Name("my_field").
		With("meta", "else").
		Fail()

	translation := translate.Simple(":attribute from :code is :value must be something @meta")(language.English, result)

	require.Equal(t, "my_field from simple is 10 must be something else", translation.Message)
}

// TestSimpleReplacesOverlappingMetadata verifies deterministic metadata replacement.
func TestSimpleReplacesOverlappingMetadata(t *testing.T) {
	result := golidate.
		Uncertain(10, "simple").
		Name("my_field").
		With("min", 2).
		With("minimum", 10).
		Fail()

	translation := translate.Simple("between @min and @minimum")(language.English, result)

	require.Equal(t, "between 2 and 10", translation.Message)
}
