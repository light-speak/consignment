package main

import (
	"context"
	"fmt"
	pb "github.com/lty5240/consignment/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
	"log"
	"os"
)

const (
	defaultHost = "mongodb://root:LTYlty0123@127.0.0.1:27017"
)

func createDummyData(repository Repository) {
	vessels := []*pb.Vessel{
		{Id: "vessel001", Name: "第一", MaxWeight: 200000, Capacity: 1100},
		{Id: "vessel002", Name: "第二", MaxWeight: 300000, Capacity: 20},
	}
	for _, v := range vessels {
		_ = repository.Create(v)
	}
}

func main() {
	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	srv.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(uri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.TODO())

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &VesselRepository{vesselCollection}
	createDummyData(repository)
	handler := &handler{repository}

	pb.RegisterVesselServiceHandler(srv.Server(), handler)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
