package main

// 1. 需要监听
// 2. 需要实例化gRPC服务端
// 3. 在gRPC注册微服务
// 4. 启动服务端

import (
	"context"
	pb "day17/04_gRPC/proto"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 地址
	addr := "127.0.0.1:8080"
	// 1. 监听
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("监听错误", err)
		return
	}
	fmt.Println("监听端口正常")
	// 2. 实例化grpc
	gs := grpc.NewServer()
	// 3. 在gRPC注册微服务
	pb.RegisterUserInfoServiceServer(gs, &u)
	// 4. 启动服务端
	gs.Serve(lis)
}

// 定义空的结构体
type UserInfoService struct {
}

var u = UserInfoService{}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, in *pb.UserRequest) (resp *pb.UserResponse, err error) {
	// 通过用户名查询用户信息
	name := in.Name
	// 数据库里查用户信息
	if name == "xyh" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   23,
			Hobby: []string{"1", "2"},
		}
	} else {
		err = errors.New("select error")
	}
	return
}
