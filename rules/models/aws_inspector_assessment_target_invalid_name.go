// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsInspectorAssessmentTargetInvalidNameRule checks the pattern is valid
type AwsInspectorAssessmentTargetInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsInspectorAssessmentTargetInvalidNameRule returns new rule with default attributes
func NewAwsInspectorAssessmentTargetInvalidNameRule() *AwsInspectorAssessmentTargetInvalidNameRule {
	return &AwsInspectorAssessmentTargetInvalidNameRule{
		resourceType:  "aws_inspector_assessment_target",
		attributeName: "name",
		max:           140,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Name() string {
	return "aws_inspector_assessment_target_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsInspectorAssessmentTargetInvalidNameRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"name must be 140 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"name must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
