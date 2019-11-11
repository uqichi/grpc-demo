package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"uqichi/grpc-demo/proto"

	"google.golang.org/grpc"

	"github.com/gorilla/mux"
)

const (
	defaultPort     = "8888"
	defaultHostAddr = "localhost:5555"
)

func Start() {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = defaultPort
	}
	hostAddr := os.Getenv("GRPC_HOST_ADDR")
	if hostAddr == "" {
		hostAddr = defaultHostAddr
	}

	// gRPC Client
	conn, err := grpc.Dial(hostAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	fmt.Println(conn.GetState())
	cli := proto.NewDemoServiceClient(conn)

	h := &handler{cli}

	r := mux.NewRouter()
	r.HandleFunc("/ping", h.pingHandler).Methods("GET")
	r.HandleFunc("/users/{uid}", h.getUserHandler).Methods("GET")
	r.HandleFunc("/users", h.createUserHandler).Methods("POST")

	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: r,
	}
	fmt.Printf("start client server on :%s\n", port)
	log.Fatal(s.ListenAndServe())
}
