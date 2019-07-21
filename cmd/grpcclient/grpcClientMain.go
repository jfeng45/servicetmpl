package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	uspb "github.com/jfeng45/servicetmpl/adapter/userclient/generatedclient"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	TARGET string = "localhost:5052"
)

func callRegisterUser(usc uspb.UserServiceClient) {
	ctx := context.Background()

	created := ptypes.TimestampNow()
	u := uspb.User{Name: "Tony", Department: "IT", Created: created}

	resp, err := usc.RegisterUser(ctx, &uspb.RegisterUserReq{User: &u})

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("results user: ", resp.User)
	}
}
func callListUser(usc uspb.UserServiceClient) {

	resp, err := usc.ListUser(context.Background(), &uspb.ListUserReq{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Got list users %s\n", resp.User)
	}
}

func main() {

	conn, err := grpc.Dial(TARGET, grpc.WithInsecure())
	if err != nil {
		fmt.Errorf("failed to dial server: %v", err)
	}
	userServiceClient := uspb.NewUserServiceClient(conn)
	fmt.Println("client strated")

	callRegisterUser(userServiceClient)
	callListUser(userServiceClient)
}
