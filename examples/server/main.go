package main

import (
	userpb "buf/example/gen/user/v1"
	"context"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

type UserService struct {
	userpb.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, req *userpb.User) (*emptypb.Empty, error) {
	if err := protovalidate.Validate(req); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	server := &UserService{}
	userpb.RegisterUserServiceServer(grpcServer, server)
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
