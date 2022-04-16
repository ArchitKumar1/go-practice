package gRPC

import (
	"context"
	"fmt"
	proto "go-practice/gRPC/protogen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GRPC struct {
	greeterClient proto.GreeterClient
}

type IgRPC interface {
	CreateServer()
	RunServer()
	sayHello(request *proto.HelloRequest) *proto.HelloReply
}

func (g *GRPC) CreateServer() {
	conn, err := grpc.Dial("localhost:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		panic("error")
	}
	g.greeterClient = proto.NewGreeterClient(conn)
}

func (g *GRPC) sayHello(request *proto.HelloRequest) *proto.HelloReply {
	ctx := context.Background()
	res, err := g.greeterClient.SayHello(ctx, request)
	if err != nil {
		fmt.Println(err)
		panic("Error")
	}
	return res
}
func (g *GRPC) RunServer() {

	for i := 0; i < 2000; i++ {
		time.Sleep(0 * time.Nanosecond)
		name := "archit"
		age := i + 1
		res, err := g.greeterClient.SayHello(context.Background(), getRequest(name, int32(age)))
		if err != nil {
			fmt.Println(err)
			panic("error")
		}
		fmt.Println(res)
	}
}

func getRequest(name string, age int32) *proto.HelloRequest {
	return &proto.HelloRequest{
		Age:  age,
		Name: name,
	}
}
