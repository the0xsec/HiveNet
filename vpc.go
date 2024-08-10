package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createVpc(ctx *pulumi.Context) (*ec2.Vpc, error) {
	return ec2.NewVpc(ctx, "hive_net", &ec2.VpcArgs{
		CidrBlock:        pulumi.String("10.0.0.0/24"),
		InstanceTenancy:  pulumi.String("default"),
		EnableDnsSupport: pulumi.Bool(true),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("hive_net"),
		},
	})
}
