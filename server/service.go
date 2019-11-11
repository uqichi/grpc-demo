package main

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"
	"uqichi/grpc-demo/proto"

	"github.com/google/uuid"

	"github.com/golang/protobuf/ptypes"
)

type demoService struct{}

type house string

const (
	Gryffindor house = "gryffindor"
	Hufflepuff house = "hufflepuff"
	Ravenclaw  house = "ravenclaw"
	Slytherin  house = "slytherin"
)

var houses = map[int]house{
	0: Gryffindor,
	1: Hufflepuff,
	2: Ravenclaw,
	3: Slytherin,
}

var mdb = sync.Map{}

type user struct {
	id      string
	name    string
	house   house
	created time.Time
}

func newDemoService() *demoService {
	return &demoService{}
}

func (svc demoService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.UserResponse, error) {
	load, ok := mdb.Load(req.Id)
	if !ok {
		return nil, errors.New("user not found")
	}
	u := load.(*user)
	ts, err := ptypes.TimestampProto(u.created)
	if err != nil {
		return nil, err
	}
	res := &proto.UserResponse{
		Id:      u.id,
		Name:    u.name,
		House:   string(u.house),
		Created: ts,
	}
	return res, nil
}

func (svc demoService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.UserResponse, error) {
	genID := uuid.New()
	userHouse := randomHouse()
	if req.House != "" {
		userHouse = house(req.House)
	}
	u := &user{
		id:      genID.String(),
		name:    req.Name,
		house:   userHouse,
		created: time.Now(),
	}
	_, loaded := mdb.LoadOrStore(genID.String(), u)
	if loaded {
		return nil, errors.New("user id already exists")
	}
	t, err := ptypes.TimestampProto(u.created)
	if err != nil {
		return nil, err
	}
	res := &proto.UserResponse{
		Id:      u.id,
		Name:    u.name,
		House:   string(u.house),
		Created: t,
	}
	return res, nil
}

func randomHouse() house {
	l := len(houses)
	return houses[rand.Int()%l]
}
