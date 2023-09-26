package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/rule"
)

func TestValue(t *testing.T) {
	value := golidate.Value(42)

	require.NotEmpty(t, value)
	require.IsType(t, golidate.Pending{}, value)
}

func TestPendingRules(t *testing.T) {
	result := golidate.
		Value(20).
		Rules(rule.Min(25)).
		Validate()

	require.False(t, result.PassesAll())
}

func TestPendingValidate(t *testing.T) {
	result := golidate.
		Value(20).
		Validate()

	require.True(t, result.PassesAll())
}
