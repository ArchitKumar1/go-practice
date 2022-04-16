package main

import (
	"go-practice/gRPC"
)

func main() {
	RunGRPCServer()

}

func RunGRPCServer() {
	var grpc gRPC.IgRPC
	grpc = &gRPC.GRPC{}
	grpc.CreateServer()
	grpc.RunServer()
}
