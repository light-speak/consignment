/*
 * @Author: Linty
 * @Github: https://github.com/lty5240
 * @Date: 2019-11-06 16:03:30
 * @LastEditor: Linty
 * @LastEditTime: 2019-11-06 17:50:40
 * @Description: -
 */
// shippy-service-consignment/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	// Import the generated protobuf code
	pb "github.com/lty5240/consignment/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/lty5240/consignment/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
)

const (
	defaultHost = "mongodb://root:LTYlty0123@127.0.0.1:27017"
)

func main() {

	srv := micro.NewService(
		micro.Name("shippy.service.consignment"),
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

	consignmentCollection := client.Database("shippy").Collection("consignments")

	repository := &MongoRepository{consignmentCollection}
	vesselClient := vesselProto.NewVesselServiceClient("shippy.service.vessel", srv.Client())
	handler := &handler{repository, vesselClient}

	pb.RegisterShippingServiceHandler(srv.Server(), handler)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
