package server

import (
	"fmt"
	"log"
	"net"
	"os"
	"uqichi/grpc-demo/proto"

	"google.golang.org/grpc"
)

func Start() {
	port := os.Getenv("GRPC_PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterDemoServiceServer(grpcServer, newDemoService())
	fmt.Printf("Start gRPC Server on :%s\n", port)
	log.Print(grpcServer.Serve(lis))
}
