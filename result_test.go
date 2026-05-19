package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
)

// TestResultFail verifies a failing result preserves code, value, and metadata.
func TestResultFail(t *testing.T) {
	res := golidate.Uncertain("something", "test").Fail()

	require.Equal(t, "test", res.Code)
	require.Equal(t, "something", res.Value)
	require.Empty(t, res.Metadata)
}

// TestResultWith verifies metadata can be added one key at a time.
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

// TestResultWithMetadata verifies metadata can be replaced wholesale.
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

// TestResultWithMetadataMerged verifies metadata maps merge without dropping keys.
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

// TestOnRenameManyDoesNotMutateSharedSlice verifies rename copies the slice.
func TestOnRenameManyDoesNotMutateSharedSlice(t *testing.T) {
	original := golidate.Results{
		golidate.Uncertain("v", "child").Name("original"),
	}

	result := golidate.
		Uncertain("v", "parent").
		With("operations", original).
		OnRename(golidate.OnRenameMany("operations"))

	_ = result.Name("renamed")

	require.Equal(t, "original", original[0].Attribute)
}

// TestResultWithMetadataMergedInitializesNilMetadata verifies nil metadata merging.
func TestResultWithMetadataMergedInitializesNilMetadata(t *testing.T) {
	result := golidate.Result{}.WithMetadataMerged(golidate.Metadata{
		"test": "foo",
	})

	require.Equal(t, golidate.Metadata{"test": "foo"}, result.Metadata)
}
