package main

// 1. 连接服务端
// 2. 实例gRPC客户端
// 3. 调用

import (
	"context"
	pb "day17/04_gRPC/proto"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	//1. 连接服务端
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		fmt.Println("连接异常")
		return
	}
	fmt.Println("连接成功")
	defer conn.Close()
	// 2. 实例gRPC客户端
	clie := pb.NewUserInfoServiceClient(conn)
	// 3. 组装一个请求参数
	req := new(pb.UserRequest)
	req.Name = "xyh"
	// 4. 调用接口
	UserResponse, err := clie.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常", err)
		return
	}
	fmt.Printf("响应结果 %v\n", UserResponse)

}
