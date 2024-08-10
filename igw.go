package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createInternetGateway(ctx *pulumi.Context, vpcID string) (*ec2.InternetGateway, error) {
	return ec2.NewInternetGateway(ctx, "hive_igw", &ec2.InternetGatewayArgs{
		VpcId: pulumi.String(vpcID),
		Tags: pulumi.StringMap{
			"Name": pulumi.String("hive_igw"),
		},
	})
}
