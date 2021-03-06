// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule checks the pattern is valid
type AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftClusterInvalidSnapshotClusterIdentifierRule returns new rule with default attributes
func NewAwsRedshiftClusterInvalidSnapshotClusterIdentifierRule() *AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule {
	return &AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule{
		resourceType:  "aws_redshift_cluster",
		attributeName: "snapshot_cluster_identifier",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule) Name() string {
	return "aws_redshift_cluster_invalid_snapshot_cluster_identifier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftClusterInvalidSnapshotClusterIdentifierRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"snapshot_cluster_identifier must be 2147483647 characters or less",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
