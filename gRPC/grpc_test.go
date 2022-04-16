package gRPC

import (
	"fmt"
	proto "go-practice/gRPC/protogen"
	"reflect"
	"testing"
)

func TestGRPC_CreateServer(t *testing.T) {
	type fields struct {
		greeterClient proto.GreeterClient
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GRPC{
				greeterClient: tt.fields.greeterClient,
			}
			fmt.Print(g)
		})
	}
}

func TestGRPC_RunServer(t *testing.T) {
	type fields struct {
		greeterClient proto.GreeterClient
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GRPC{
				greeterClient: tt.fields.greeterClient,
			}
			fmt.Print(g)
		})
	}
}

func TestGRPC_sayHello(t *testing.T) {
	type fields struct {
		greeterClient proto.GreeterClient
	}
	type args struct {
		request *proto.HelloRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *proto.HelloReply
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GRPC{
				greeterClient: tt.fields.greeterClient,
			}
			if got := g.sayHello(tt.args.request); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRequest(t *testing.T) {
	type args struct {
		name string
		age  int32
	}
	tests := []struct {
		name string
		args args
		want *proto.HelloRequest
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRequest(tt.args.name, tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

//  mockery --name=IgRPC --dir ./gRPC
