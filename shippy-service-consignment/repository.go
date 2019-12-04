package main

import (
	"context"
	pb "github.com/lty5240/consignment/shippy-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type repository interface {
	Create(consignment *pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
}

type MongoRepository struct {
	collection *mongo.Collection
}

func (repository *MongoRepository) Create(consignment *pb.Consignment) error {
	_, err := repository.collection.InsertOne(context.Background(), consignment)
	return err
}

func (repository *MongoRepository) GetAll() ([]*pb.Consignment, error) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := repository.collection.Find(ctx, bson.D{})
	var consignments []*pb.Consignment
	if cur == nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var consignment *pb.Consignment
		if err := cur.Decode(&consignment); err != nil {
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
