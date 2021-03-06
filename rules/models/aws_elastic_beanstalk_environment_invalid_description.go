// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsElasticBeanstalkEnvironmentInvalidDescriptionRule checks the pattern is valid
type AwsElasticBeanstalkEnvironmentInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsElasticBeanstalkEnvironmentInvalidDescriptionRule returns new rule with default attributes
func NewAwsElasticBeanstalkEnvironmentInvalidDescriptionRule() *AwsElasticBeanstalkEnvironmentInvalidDescriptionRule {
	return &AwsElasticBeanstalkEnvironmentInvalidDescriptionRule{
		resourceType:  "aws_elastic_beanstalk_environment",
		attributeName: "description",
		max:           200,
	}
}

// Name returns the rule name
func (r *AwsElasticBeanstalkEnvironmentInvalidDescriptionRule) Name() string {
	return "aws_elastic_beanstalk_environment_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElasticBeanstalkEnvironmentInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElasticBeanstalkEnvironmentInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElasticBeanstalkEnvironmentInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElasticBeanstalkEnvironmentInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 200 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
