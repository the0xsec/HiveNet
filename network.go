package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createNetworkAcl(ctx *pulumi.Context, vpcID, myIPCidr string) (*ec2.NetworkAcl, error) {
	return ec2.NewNetworkAcl(ctx, "hive_egress_nacl", &ec2.NetworkAclArgs{
		VpcId: pulumi.String(vpcID),
		Egress: ec2.NetworkAclEgressArray{
			&ec2.NetworkAclEgressArgs{
				Protocol:  pulumi.String("-1"),
				RuleNo:    pulumi.Int(100),
				Action:    pulumi.String("allow"),
				CidrBlock: pulumi.String("0.0.0.0/0"),
				FromPort:  pulumi.Int(0),
				ToPort:    pulumi.Int(0),
			},
		},
		Ingress: ec2.NetworkAclIngressArray{
			&ec2.NetworkAclIngressArgs{
				Protocol:  pulumi.String("tcp"),
				RuleNo:    pulumi.Int(100),
				Action:    pulumi.String("allow"),
				CidrBlock: pulumi.String(myIPCidr),
				FromPort:  pulumi.Int(80),
				ToPort:    pulumi.Int(80),
			},
			&ec2.NetworkAclIngressArgs{
				Protocol:  pulumi.String("tcp"),
				RuleNo:    pulumi.Int(101),
				Action:    pulumi.String("allow"),
				CidrBlock: pulumi.String(myIPCidr),
				FromPort:  pulumi.Int(443),
				ToPort:    pulumi.Int(443),
			},
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("hive_egress_nacl"),
		},
	})
}
