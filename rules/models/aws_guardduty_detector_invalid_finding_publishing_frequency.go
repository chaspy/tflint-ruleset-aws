// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule checks the pattern is valid
type AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule returns new rule with default attributes
func NewAwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule() *AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule {
	return &AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule{
		resourceType:  "aws_guardduty_detector",
		attributeName: "finding_publishing_frequency",
		enum: []string{
			"FIFTEEN_MINUTES",
			"ONE_HOUR",
			"SIX_HOURS",
		},
	}
}

// Name returns the rule name
func (r *AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule) Name() string {
	return "aws_guardduty_detector_invalid_finding_publishing_frequency"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsGuarddutyDetectorInvalidFindingPublishingFrequencyRule) Check(runner tflint.Runner) error {
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
					fmt.Sprintf(`"%s" is an invalid value as finding_publishing_frequency`, truncateLongMessage(val)),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
