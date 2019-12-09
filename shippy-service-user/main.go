package main

import (
	"fmt"
	pb "github.com/lty5240/consignment/shippy-service-user/proto/user"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	db, err := CreateConnection()
	if db == nil {
		log.Panic("db not nil")
	}
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	db.AutoMigrate(&pb.User{})

	repository := &UserRepository{db}
	tokenService := &TokenService{repository}

	srv := micro.NewService(
		micro.Name("shippy.service.user"),
		micro.Version("latest"),
	)
	srv.Init()
	pb.RegisterUserServiceHandler(srv.Server(), &service{repository, tokenService})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
