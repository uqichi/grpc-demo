package main

import (
	"fmt"
	"google.golang.org/grpc/credentials"
	"log"
	"net/http"
	"os"
	"uqichi/grpc-demo/proto"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
)

const (
	defaultPort         = "8888"
	defaultGRPCHostAddr = "localhost:5555"
	defaultHTTPHostAddr = "localhost:8000"
)

func Start() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	}
	gRPCHost := os.Getenv("GRPC_HOST_ADDR")
	if gRPCHost == "" {
		gRPCHost = defaultGRPCHostAddr
	}
	httpHost := os.Getenv("HTTP_HOST_ADDR")
	if httpHost == "" {
		httpHost = defaultHTTPHostAddr
	}

	// gRPC Client
	creds, err := credentials.NewClientTLSFromFile("/tls/ca.crt", "test.jp")
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}
	conn, err := grpc.Dial(gRPCHost, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	fmt.Println(conn.GetState())
	cli := proto.NewDemoServiceClient(conn)

	h := &handler{cli, httpHost}

	r := mux.NewRouter()
	r.HandleFunc("/httpping", h.httpPingHandler).Methods("GET")
	r.HandleFunc("/grpcping", h.grpcPingHandler).Methods("GET")
	r.HandleFunc("/users/{uid}", h.getUserHandler).Methods("GET")
	r.HandleFunc("/users", h.listUsersHandler).Methods("GET")
	r.HandleFunc("/users", h.createUserHandler).Methods("POST")

	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: r,
	}
	fmt.Printf("start client server on :%s\n", port)
	log.Fatal(s.ListenAndServe())
}
