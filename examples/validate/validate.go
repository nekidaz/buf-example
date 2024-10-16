package main

import (
	userpb "buf/example/gen/user/v1"
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"log"
)

func main() {
	user := &userpb.User{
		Name:     "Dilmurat",
		Email:    "dlol96@mail.ru",
		Password: "Kosynvoa110!",
	}

	if err := protovalidate.Validate(user); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("user: %v\n", user)

}
