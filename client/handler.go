package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"uqichi/grpc-demo/proto"

	"github.com/golang/protobuf/ptypes"

	"github.com/gorilla/mux"
)

type User struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	House   string    `json:"house"`
	Created time.Time `json:"created"`
}

type handler struct {
	cli proto.DemoServiceClient
}

func (h *handler) getUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	res, err := h.cli.GetUser(r.Context(), &proto.GetUserRequest{Id: vars["uid"]})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	ts, err := ptypes.Timestamp(res.Created)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	user := User{
		Name:    res.Name,
		Created: ts,
	}

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func (h *handler) createUserHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	res, err := h.cli.CreateUser(r.Context(), &proto.CreateUserRequest{
		Name:  user.Name,
		House: user.House,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Response: %v\n", res)
}
