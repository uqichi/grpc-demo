package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"uqichi/grpc-demo/proto"

	"google.golang.org/grpc"
)

const (
	defaultPort = "5555"
)

func Start() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = defaultPort
	}

	// gPRC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterDemoServiceServer(grpcServer, newDemoService())

	fmt.Printf("start server on :%s\n", port)
	log.Fatal(grpcServer.Serve(lis))
}
