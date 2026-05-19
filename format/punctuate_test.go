package format_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate/format"
)

// TestPunctuateAppendsPeriod verifies a period is added when missing.
func TestPunctuateAppendsPeriod(t *testing.T) {
	formatter := format.Punctuate()

	require.Equal(t, "hello.", formatter("hello"))
}

// TestPunctuateSkipsExistingPeriod verifies no double period.
func TestPunctuateSkipsExistingPeriod(t *testing.T) {
	formatter := format.Punctuate()

	require.Equal(t, "hello.", formatter("hello."))
}

// TestPunctuateSkipsExclamation verifies exclamation is not followed by period.
func TestPunctuateSkipsExclamation(t *testing.T) {
	formatter := format.Punctuate()

	require.Equal(t, "hello!", formatter("hello!"))
}

// TestPunctuateSkipsQuestion verifies question mark is not followed by period.
func TestPunctuateSkipsQuestion(t *testing.T) {
	formatter := format.Punctuate()

	require.Equal(t, "hello?", formatter("hello?"))
}

// TestPunctuateEmptyString verifies empty input is returned unchanged.
func TestPunctuateEmptyString(t *testing.T) {
	formatter := format.Punctuate()

	require.Equal(t, "", formatter(""))
}
