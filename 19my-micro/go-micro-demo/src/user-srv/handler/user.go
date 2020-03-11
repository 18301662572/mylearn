package handler

import (
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/share/pb"
	tlog "code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/share/utils/log"
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/user-srv/db"
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/user-srv/entity"
	"context"
	"github.com/golang/protobuf/go/src/pkg/log"
	"go.uber.org/zap"
)

//BLL:业务逻辑层

//定义绑定方法的结构体
type UserHandler struct {
	logger *zap.Logger
}

//创建结构体对象
func NewUserHandler() *UserHandler {
	return &UserHandler{logger: tlog.Instance().Named("UserHandler")}
}

func (h *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserReq, resp *pb.InsertUserResp) error {
	log.Println("InsertUser...")
	//封装结构体
	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}
	//调用数据库的方法进行插入
	insertId, err := db.InsertUser(user)
	if err != nil {
		log.Fatal("添加用户错误")
		return err
	}
	resp.Id = int32(insertId)
	return nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserReq, resp *pb.DeleteUserResp) error {
	log.Println("DeleteUser...")
	//调用数据库的方法删除
	err := db.DeleteUser(req.Id)
	if err != nil {
		log.Fatal("删除用户错误")
		return err
	}
	return nil
}

func (h *UserHandler) ModifyUser(ctx context.Context, req *pb.ModifyUserReq, resp *pb.ModifyUserResp) error {
	log.Println("ModifyUser...", req.GetId())
	//封装结构体
	user := &entity.User{
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
		Id:      req.Id,
	}
	//调用数据库的方法进行修改
	err := db.ModifyUser(user)
	if err != nil {
		log.Fatal("修改用户错误")
		return err
	}
	return nil
}

func (h *UserHandler) SelectUser(ctx context.Context, req *pb.SelectUserReq, resp *pb.SelectUserResp) error {
	log.Println("SelectUser...")
	//调用数据库的方法进行查询
	user, err := db.SelectUserById(req.GetId())
	if err != nil {
		log.Fatal("查询用户错误")
		return err
	}
	if user != nil {
		resp.Users = user.ToProtoUser()
	}
	return nil
}
