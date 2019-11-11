package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	defaultPort = "8888"
)

func Start() {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = defaultPort
	}
	r := mux.NewRouter()
	r.HandleFunc("/ping", pingHandler).Methods("GET")
	r.HandleFunc("/users/{uid}", getUserHandler).Methods("GET")
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		Handler: r,
	}
	fmt.Printf("Start gRPC Client on :%s\n", port)
	log.Fatal(s.ListenAndServe())
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Pong")
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "uid: %v\n", vars["uid"])

	// TODO: get

	var user User
	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// TODO: create

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
