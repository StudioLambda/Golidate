package golidate_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/studiolambda/golidate"
	"github.com/studiolambda/golidate/format"
)

func TestGroupedMessages(t *testing.T) {
	grouped := golidate.Grouped{
		"name": golidate.Results{
			golidate.Fail("", "required").Name("name").WithMetadata(golidate.Metadata{}),
		},
		"email": golidate.Results{
			golidate.Fail("invalid", "email").Name("email").WithMetadata(golidate.Metadata{}),
		},
	}
	grouped["name"][0].Message = "name is required"
	grouped["email"][0].Message = "email is invalid"

	messages := grouped.Messages(format.Capitalize(), format.Punctuate())

	require.Equal(t, []string{"Name is required."}, messages["name"])
	require.Equal(t, []string{"Email is invalid."}, messages["email"])
}
