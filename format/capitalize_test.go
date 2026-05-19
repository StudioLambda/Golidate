package format_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate/format"
)

// TestCapitalizeHandlesUnicode verifies capitalization handles multibyte runes.
func TestCapitalizeHandlesUnicode(t *testing.T) {
	formatter := format.Capitalize()

	require.Equal(t, "Äpfel", formatter("äpfel"))
}

// TestUncapitalizeHandlesUnicode verifies uncapitalization handles multibyte runes.
func TestUncapitalizeHandlesUnicode(t *testing.T) {
	formatter := format.Uncapitalize()

	require.Equal(t, "äpfel", formatter("Äpfel"))
}
