package main

import (
	"fmt"
	pb "gin-test/grpc/proto" // 引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = ":50053"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	rsp, err := c.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Message)
}