package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func createRouteTable(ctx *pulumi.Context, vpcID, gatewayID string) (*ec2.RouteTable, error) {
	return ec2.NewRouteTable(ctx, "hive_route_table", &ec2.RouteTableArgs{
		VpcId: pulumi.String(vpcID),
		Routes: ec2.RouteTableRouteArray{
			&ec2.RouteTableRouteArgs{
				CidrBlock: pulumi.String("0.0.0.0/0"),
				GatewayId: pulumi.String(gatewayID),
			},
		},
		Tags: pulumi.StringMap{
			"Name": pulumi.String("hive_route_table"),
		},
	})
}

func associateRouteTable(ctx *pulumi.Context, subnetID, routeTableID string) error {
	_, err := ec2.NewRouteTableAssociation(ctx, "hive_route_table_association", &ec2.RouteTableAssociationArgs{
		SubnetId:     pulumi.String(subnetID),
		RouteTableId: pulumi.String(routeTableID),
	})
	return err
}
