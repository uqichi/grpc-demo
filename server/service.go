package server

import (
	"context"
	"errors"
	"sync"
	"time"
	"uqichi/grpc-demo/proto"

	"github.com/google/uuid"

	"github.com/golang/protobuf/ptypes"
)

type demoService struct {
	m sync.Map
}

type user struct {
	id      string
	name    string
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
		Created: ts,
	}
	return res, nil
}

func (svc demoService) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
	genID := uuid.New()
	u := &user{
		id:      genID.String(),
		name:    req.Name,
		created: time.Now(),
	}
	_, loaded := svc.m.LoadOrStore(genID.String(), u)
	if loaded {
		return nil, errors.New("user id already exists")
	}
	t, err := ptypes.TimestampProto(u.created)
	if err == nil {
		return nil, err
	}
	res := &proto.CreateUserResponse{
		Id:      u.id,
		Created: t,
	}
	return res, nil
}
