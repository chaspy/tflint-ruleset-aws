// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCloudwatchEventPermissionInvalidPrincipalRule checks the pattern is valid
type AwsCloudwatchEventPermissionInvalidPrincipalRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCloudwatchEventPermissionInvalidPrincipalRule returns new rule with default attributes
func NewAwsCloudwatchEventPermissionInvalidPrincipalRule() *AwsCloudwatchEventPermissionInvalidPrincipalRule {
	return &AwsCloudwatchEventPermissionInvalidPrincipalRule{
		resourceType:  "aws_cloudwatch_event_permission",
		attributeName: "principal",
		max:           12,
		min:           1,
		pattern:       regexp.MustCompile(`^(\d{12}|\*)$`),
	}
}

// Name returns the rule name
func (r *AwsCloudwatchEventPermissionInvalidPrincipalRule) Name() string {
	return "aws_cloudwatch_event_permission_invalid_principal"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchEventPermissionInvalidPrincipalRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudwatchEventPermissionInvalidPrincipalRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchEventPermissionInvalidPrincipalRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchEventPermissionInvalidPrincipalRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"principal must be 12 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"principal must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^(\d{12}|\*)$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}