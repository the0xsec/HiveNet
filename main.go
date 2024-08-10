package main

import (
	"io"
	"net/http"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func getPublicIP() (string, string, error) {
	resp, err := http.Get("http://checkip.amazonaws.com/")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	myIP := strings.TrimSpace(string(body))
	myIPCidr := myIP + "/32"
	return myIP, myIPCidr, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpc, err := createVpc(ctx)
		if err != nil {
			return err
		}

		myIP, myIPCidr, err := getPublicIP()
		if err != nil {
			return err
		}

		subnet1, err := createSubnet(ctx, "hive_subnet", vpc.ID(), "10.0.0.0/26", "us-east-1a")
		if err != nil {
			return err
		}

		_, err = createSubnet(ctx, "hive_subnet2", vpc.ID(), "10.0.0.64/26", "us-east-1b")
		if err != nil {
			return err
		}

		_, err = createSubnet(ctx, "hive_subnet3", vpc.ID(), "10.0.0.128/26", "us-east-1c")
		if err != nil {
			return err
		}

		_, err = createNetworkAcl(ctx, vpc.ID(), myIPCidr)
		if err != nil {
			return err
		}

		igw, err := createInternetGateway(ctx, vpc.ID())
		if err != nil {
			return err
		}

		routeTable, err := createRouteTable(ctx, vpc.ID(), igw.ID())
		if err != nil {
			return err
		}

		err = associateRouteTable(ctx, subnet1.ID(), routeTable.ID())
		if err != nil {
			return err
		}

		exportOutputs(ctx, vpc.ID(), myIP)

		return nil
	})
}
