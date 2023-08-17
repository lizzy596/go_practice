package main

import (
	"fmt"
	"log"

	pb "grpc-client.com/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	fmt.Println(". connect to gRPC addr [", addr, "]")
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf(">>> err: connect failed: %v\n", err)
	}

	defer conn.Close()

	fmt.Println("> create new  gRPC client ...")
	c := pb.NewControlPlaneClient(conn)

	sendRequest(c)
}
