package main

import (
	"context"
	"fmt"
	"log"

	pb "grpc-client.com/proto/control-plane"
)

func sendRequest(c pb.ControlPlaneClient) {

	var device_name = "gps"
	var command = "on"

	fmt.Println(">> send gRPC Request:")
	fmt.Println("  - device_name:", device_name)
	fmt.Println("  - command:", command)

	res, err := c.DeviceRequest(context.Background(), &pb.Options{
		DeviceName: device_name,
		Command:    command,
	})

	if err != nil {
		log.Fatalf(">>> err: gRPC request failed: %v\n", err)
	}

	fmt.Println("<<< recv gRPC Response: [", res.Response, "]")
}
