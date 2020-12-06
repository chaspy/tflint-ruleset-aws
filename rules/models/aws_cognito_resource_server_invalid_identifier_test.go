// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsCognitoResourceServerInvalidIdentifierRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cognito_resource_server" "foo" {
	identifier = "	"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsCognitoResourceServerInvalidIdentifierRule(),
					Message: `"	" does not match valid pattern ^[\x21\x23-\x5B\x5D-\x7E]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cognito_resource_server" "foo" {
	identifier = "https://example.com"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsCognitoResourceServerInvalidIdentifierRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}