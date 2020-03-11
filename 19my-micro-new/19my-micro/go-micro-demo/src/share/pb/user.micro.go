// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserService interface {
	//增删改查
	InsertUser(ctx context.Context, in *InsertUserReq, opts ...client.CallOption) (*InsertUserResp, error)
	DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...client.CallOption) (*DeleteUserResp, error)
	ModifyUser(ctx context.Context, in *ModifyUserReq, opts ...client.CallOption) (*ModifyUserResp, error)
	SelectUser(ctx context.Context, in *SelectUserReq, opts ...client.CallOption) (*SelectUserResp, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "pb"
	}
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) InsertUser(ctx context.Context, in *InsertUserReq, opts ...client.CallOption) (*InsertUserResp, error) {
	req := c.c.NewRequest(c.name, "UserService.InsertUser", in)
	out := new(InsertUserResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) DeleteUser(ctx context.Context, in *DeleteUserReq, opts ...client.CallOption) (*DeleteUserResp, error) {
	req := c.c.NewRequest(c.name, "UserService.DeleteUser", in)
	out := new(DeleteUserResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ModifyUser(ctx context.Context, in *ModifyUserReq, opts ...client.CallOption) (*ModifyUserResp, error) {
	req := c.c.NewRequest(c.name, "UserService.ModifyUser", in)
	out := new(ModifyUserResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) SelectUser(ctx context.Context, in *SelectUserReq, opts ...client.CallOption) (*SelectUserResp, error) {
	req := c.c.NewRequest(c.name, "UserService.SelectUser", in)
	out := new(SelectUserResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	//增删改查
	InsertUser(context.Context, *InsertUserReq, *InsertUserResp) error
	DeleteUser(context.Context, *DeleteUserReq, *DeleteUserResp) error
	ModifyUser(context.Context, *ModifyUserReq, *ModifyUserResp) error
	SelectUser(context.Context, *SelectUserReq, *SelectUserResp) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		InsertUser(ctx context.Context, in *InsertUserReq, out *InsertUserResp) error
		DeleteUser(ctx context.Context, in *DeleteUserReq, out *DeleteUserResp) error
		ModifyUser(ctx context.Context, in *ModifyUserReq, out *ModifyUserResp) error
		SelectUser(ctx context.Context, in *SelectUserReq, out *SelectUserResp) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) InsertUser(ctx context.Context, in *InsertUserReq, out *InsertUserResp) error {
	return h.UserServiceHandler.InsertUser(ctx, in, out)
}

func (h *userServiceHandler) DeleteUser(ctx context.Context, in *DeleteUserReq, out *DeleteUserResp) error {
	return h.UserServiceHandler.DeleteUser(ctx, in, out)
}

func (h *userServiceHandler) ModifyUser(ctx context.Context, in *ModifyUserReq, out *ModifyUserResp) error {
	return h.UserServiceHandler.ModifyUser(ctx, in, out)
}

func (h *userServiceHandler) SelectUser(ctx context.Context, in *SelectUserReq, out *SelectUserResp) error {
	return h.UserServiceHandler.SelectUser(ctx, in, out)
}
