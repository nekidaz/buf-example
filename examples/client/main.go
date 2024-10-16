package main

import (
	userpb "buf/example/gen/user/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)
	user := &userpb.User{
		Name:     "Dilmurat",
		Email:    "",
		Password: "Kosynvoa110!",
	}
	_, err = client.CreateUser(context.Background(), user)
	if err != nil {
		panic(err)
	}
	fmt.Println("User created")

}
