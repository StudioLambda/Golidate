package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
)

func TestResultFail(t *testing.T) {
	res := golidate.Uncertain("something", "test").Fail()

	require.Equal(t, "test", res.Code)
	require.Equal(t, "something", res.Value)
	require.Empty(t, res.Metadata)
}

func TestResultWith(t *testing.T) {
	res := golidate.
		Uncertain("something", "test").
		With("test", "foo").
		With("test2", "bar").
		Fail()

	require.Equal(t, "test", res.Code)
	require.Equal(t, "something", res.Value)
	require.NotEmpty(t, res.Metadata)
	require.Equal(t, golidate.Metadata{"test": "foo", "test2": "bar"}, res.Metadata)
}

func TestResultWithMetadata(t *testing.T) {
	metadata := golidate.Metadata{
		"test": "foo",
	}

	res := golidate.
		Uncertain("something", "test").
		WithMetadata(metadata).
		Fail()

	require.Equal(t, "test", res.Code)
	require.Equal(t, "something", res.Value)
	require.NotEmpty(t, res.Metadata)
	require.Equal(t, metadata, res.Metadata)
}

func TestResultWithMetadataMerged(t *testing.T) {
	metadata := golidate.Metadata{
		"test": "foo",
	}

	metadata2 := golidate.Metadata{
		"test2": "foo",
	}

	res := golidate.
		Uncertain("something", "test").
		WithMetadata(metadata).
		WithMetadataMerged(metadata2).
		Fail()

	require.Equal(t, "test", res.Code)
	require.Equal(t, "something", res.Value)
	require.NotEmpty(t, res.Metadata)
	require.Equal(t, golidate.Metadata{"test": "foo", "test2": "foo"}, res.Metadata)
}
