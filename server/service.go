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

type demoService struct {
	m sync.Map
}

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

type user struct {
	id      string
	name    string
	house   house
	created time.Time
}

func newDemoService() *demoService {
	return &demoService{m: sync.Map{}}
}

func (svc demoService) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	load, ok := svc.m.Load(req.Id)
	if !ok {
		return nil, errors.New("user not found")
	}
	u := load.(*user)
	ts, err := ptypes.TimestampProto(u.created)
	if err != nil {
		return nil, err
	}
	res := &proto.GetUserResponse{
		Name:    u.name,
		House:   string(u.house),
		Created: ts,
	}
	return res, nil
}

func (svc demoService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
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
	_, loaded := svc.m.LoadOrStore(genID.String(), u)
	if loaded {
		return nil, errors.New("user id already exists")
	}
	t, err := ptypes.TimestampProto(u.created)
	if err != nil {
		return nil, err
	}
	res := &proto.CreateUserResponse{
		Id:      u.id,
		House:   string(u.house),
		Created: t,
	}
	return res, nil
}

func randomHouse() house {
	l := len(houses)
	return houses[rand.Int()%l]
}
