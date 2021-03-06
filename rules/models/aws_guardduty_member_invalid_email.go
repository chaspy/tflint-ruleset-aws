// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyMemberInvalidEmailRule checks the pattern is valid
type AwsGuarddutyMemberInvalidEmailRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsGuarddutyMemberInvalidEmailRule returns new rule with default attributes
func NewAwsGuarddutyMemberInvalidEmailRule() *AwsGuarddutyMemberInvalidEmailRule {
	return &AwsGuarddutyMemberInvalidEmailRule{
		resourceType:  "aws_guardduty_member",
		attributeName: "email",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsGuarddutyMemberInvalidEmailRule) Name() string {
	return "aws_guardduty_member_invalid_email"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyMemberInvalidEmailRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyMemberInvalidEmailRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyMemberInvalidEmailRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyMemberInvalidEmailRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"email must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"email must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
