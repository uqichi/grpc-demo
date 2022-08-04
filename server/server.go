package main

import (
	"fmt"
	"google.golang.org/grpc/credentials"
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

	creds, err := credentials.NewServerTLSFromFile("/tls/service.pem", "/tls/service.key")
	if err != nil {
		log.Fatalf("failed to create credentials: %v", err)
	}

	// gPRC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	proto.RegisterDemoServiceServer(grpcServer, newDemoService())

	fmt.Printf("start grpc server on :%s\n", port)
	log.Fatal(grpcServer.Serve(lis))
}

func StartHTTP() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultHTTPPort
	}

	// http Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("http ping", os.Getenv("MY_POD_IP"), r.RemoteAddr)
		fmt.Println(r.Header.Get("X-Forwarded-For"))
		fmt.Fprintf(w, "podip - %s, remoteAddr - %s", os.Getenv("MY_POD_IP"), r.RemoteAddr)
	})
	fmt.Printf("start http server on :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
