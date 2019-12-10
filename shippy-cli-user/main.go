package main

import (
	"context"
	pb "github.com/lty5240/consignment/shippy-service-user/proto/user"
	microClient "github.com/micro/go-micro/client"
	"log"
	"os"
)

func main() {
	client := pb.NewUserServiceClient("shippy.service.user", microClient.DefaultClient)

	name := "linty"
	email := "1046044814@qq.com"
	password := "lty01234"
	company := "KLM"

	log.Println(name, email, password)

	res, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("could not create : %v", err)
	}
	log.Printf("Created : %s", res.User.Id)

	getALl, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users : %v", err)
	}
	for _, v := range getALl.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user : %s error : %v\n", email, err)
	}

	log.Printf("Your assess token is %s \n", authResponse.Token)
	os.Exit(0)

}
