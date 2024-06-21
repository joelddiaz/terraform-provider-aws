// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/networkfirewall"
)

// FindFirewallPolicyByNameAndARN returns the FirewallPolicyOutput from a call to DescribeFirewallPolicyWithContext
// given the context and at least one of FirewallPolicyArn and FirewallPolicyName.
func FindFirewallPolicyByNameAndARN(ctx context.Context, conn *networkfirewall.NetworkFirewall, arn string, name string) (*networkfirewall.DescribeFirewallPolicyOutput, error) {
	input := &networkfirewall.DescribeFirewallPolicyInput{}
	if arn != "" {
		input.FirewallPolicyArn = aws.String(arn)
	}
	if name != "" {
		input.FirewallPolicyName = aws.String(name)
	}

	output, err := conn.DescribeFirewallPolicyWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// FindResourcePolicy returns the Policy string from a call to DescribeResourcePolicyWithContext
// given the context and resource ARN.
// Returns nil if the FindResourcePolicy is not found.
func FindResourcePolicy(ctx context.Context, conn *networkfirewall.NetworkFirewall, arn string) (*string, error) {
	input := &networkfirewall.DescribeResourcePolicyInput{
		ResourceArn: aws.String(arn),
	}
	output, err := conn.DescribeResourcePolicyWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, nil
	}
	return output.Policy, nil
}
