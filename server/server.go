package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"uqichi/grpc-demo/proto"

	"google.golang.org/grpc"
)

const (
	defaultGRPCPort = "5555"
	defaultHTTPPort = "8000"
)

func StartGRPC() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = defaultGRPCPort
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

func StartHTTP() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultHTTPPort
	}

	// http Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "podip - %s", os.Getenv("MY_POD_IP"))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
