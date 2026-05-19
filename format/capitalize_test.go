package format_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate/format"
)

func TestCapitalizeHandlesUnicode(t *testing.T) {
	formatter := format.Capitalize()

	require.Equal(t, "Äpfel", formatter("äpfel"))
}

func TestUncapitalizeHandlesUnicode(t *testing.T) {
	formatter := format.Uncapitalize()

	require.Equal(t, "äpfel", formatter("Äpfel"))
}
