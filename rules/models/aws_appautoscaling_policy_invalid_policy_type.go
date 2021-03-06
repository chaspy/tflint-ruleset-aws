// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsAppautoscalingPolicyInvalidPolicyTypeRule checks the pattern is valid
type AwsAppautoscalingPolicyInvalidPolicyTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsAppautoscalingPolicyInvalidPolicyTypeRule returns new rule with default attributes
func NewAwsAppautoscalingPolicyInvalidPolicyTypeRule() *AwsAppautoscalingPolicyInvalidPolicyTypeRule {
	return &AwsAppautoscalingPolicyInvalidPolicyTypeRule{
		resourceType:  "aws_appautoscaling_policy",
		attributeName: "policy_type",
		enum: []string{
			"StepScaling",
			"TargetTrackingScaling",
		},
	}
}

// Name returns the rule name
func (r *AwsAppautoscalingPolicyInvalidPolicyTypeRule) Name() string {
	return "aws_appautoscaling_policy_invalid_policy_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppautoscalingPolicyInvalidPolicyTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsAppautoscalingPolicyInvalidPolicyTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsAppautoscalingPolicyInvalidPolicyTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppautoscalingPolicyInvalidPolicyTypeRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is an invalid value as policy_type`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
