// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsConfigConfigurationAggregatorInvalidNameRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_config_configuration_aggregator" "foo" {
	name = "example.com"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewAwsConfigConfigurationAggregatorInvalidNameRule(),
					Message: `"example.com" does not match valid pattern ^[\w\-]+$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_config_configuration_aggregator" "foo" {
	name = "example"
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewAwsConfigConfigurationAggregatorInvalidNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}
