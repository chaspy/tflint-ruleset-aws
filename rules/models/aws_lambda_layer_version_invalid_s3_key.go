// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsLambdaLayerVersionInvalidS3KeyRule checks the pattern is valid
type AwsLambdaLayerVersionInvalidS3KeyRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsLambdaLayerVersionInvalidS3KeyRule returns new rule with default attributes
func NewAwsLambdaLayerVersionInvalidS3KeyRule() *AwsLambdaLayerVersionInvalidS3KeyRule {
	return &AwsLambdaLayerVersionInvalidS3KeyRule{
		resourceType:  "aws_lambda_layer_version",
		attributeName: "s3_key",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsLambdaLayerVersionInvalidS3KeyRule) Name() string {
	return "aws_lambda_layer_version_invalid_s3_key"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsLambdaLayerVersionInvalidS3KeyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsLambdaLayerVersionInvalidS3KeyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsLambdaLayerVersionInvalidS3KeyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsLambdaLayerVersionInvalidS3KeyRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"s3_key must be 1024 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"s3_key must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}