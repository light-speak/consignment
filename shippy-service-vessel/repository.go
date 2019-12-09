package main

import (
	"context"
	pb "github.com/lty5240/consignment/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(*pb.Vessel) error
}

type VesselRepository struct {
	collection *mongo.Collection
}

func (repository *VesselRepository) FindAvailable(specification *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel
	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := repository.collection.FindOne(context.TODO(), bson.M{
		"capacity":  bson.M{"$gte": specification.Capacity},
		"maxweight": bson.M{"$gte": specification.MaxWeight},
	}).Decode(&vessel)

	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repository *VesselRepository) Create(vessel *pb.Vessel) error {
	_, err := repository.collection.InsertOne(context.Background(), vessel)
	if err != nil {
		return err
	}
	return nil
}
