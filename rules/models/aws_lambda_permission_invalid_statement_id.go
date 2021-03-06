// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaPermissionInvalidStatementIDRule checks the pattern is valid
type AwsLambdaPermissionInvalidStatementIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsLambdaPermissionInvalidStatementIDRule returns new rule with default attributes
func NewAwsLambdaPermissionInvalidStatementIDRule() *AwsLambdaPermissionInvalidStatementIDRule {
	return &AwsLambdaPermissionInvalidStatementIDRule{
		resourceType:  "aws_lambda_permission",
		attributeName: "statement_id",
		max:           100,
		min:           1,
		pattern:       regexp.MustCompile(`^([a-zA-Z0-9-_]+)$`),
	}
}

// Name returns the rule name
func (r *AwsLambdaPermissionInvalidStatementIDRule) Name() string {
	return "aws_lambda_permission_invalid_statement_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaPermissionInvalidStatementIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaPermissionInvalidStatementIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaPermissionInvalidStatementIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaPermissionInvalidStatementIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"statement_id must be 100 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"statement_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^([a-zA-Z0-9-_]+)$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
