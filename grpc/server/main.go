package main

import (
	"fmt"
	"net"

	pb "gin-test/grpc/proto" // 引入编译生成的包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//tls以及 token认真 参见 http://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/gRPC/%E8%AE%A4%E8%AF%81.html

const (
	// Address gRPC服务地址
	Address = ":50053"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello jdx %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Println("Failed to listen:", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	fmt.Println("Listen on " + Address)
	s.Serve(listen)
}