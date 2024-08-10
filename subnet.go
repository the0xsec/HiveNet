package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createSubnet(ctx *pulumi.Context, name, vpcID, cidrBlock, availabilityZone string) (*ec2.Subnet, error) {
	return ec2.NewSubnet(ctx, name, &ec2.SubnetArgs{
		VpcId:            pulumi.String(vpcID),
		CidrBlock:        pulumi.String(cidrBlock),
		AvailabilityZone: pulumi.String(availabilityZone),
		Tags: pulumi.StringMap{
			"Name": pulumi.String(name),
		},
	})
}
