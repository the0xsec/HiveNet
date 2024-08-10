package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func exportOutputs(ctx *pulumi.Context, vpcID, myIP string) {
	ctx.Export("vpcId", pulumi.String(vpcID))
	ctx.Export("local_ip", pulumi.String(myIP))
}
