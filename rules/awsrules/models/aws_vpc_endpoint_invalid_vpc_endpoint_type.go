// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsVpcEndpointInvalidVpcEndpointTypeRule checks the pattern is valid
type AwsVpcEndpointInvalidVpcEndpointTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsVpcEndpointInvalidVpcEndpointTypeRule returns new rule with default attributes
func NewAwsVpcEndpointInvalidVpcEndpointTypeRule() *AwsVpcEndpointInvalidVpcEndpointTypeRule {
	return &AwsVpcEndpointInvalidVpcEndpointTypeRule{
		resourceType:  "aws_vpc_endpoint",
		attributeName: "vpc_endpoint_type",
		enum: []string{
			"Interface",
			"Gateway",
			"GatewayLoadBalancer",
		},
	}
}

// Name returns the rule name
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Name() string {
	return "aws_vpc_endpoint_invalid_vpc_endpoint_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsVpcEndpointInvalidVpcEndpointTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as vpc_endpoint_type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}
