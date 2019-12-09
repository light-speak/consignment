package main

import (
	"context"
	pb "github.com/lty5240/consignment/shippy-service-vessel/proto/vessel"
)

type handler struct {
	Repository
}

func (handler *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	vessel, err := handler.Repository.FindAvailable(req)
	if err != nil {
		return err
	}
	res.Vessel = vessel
	return nil
}

func (handler *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	err := handler.Repository.Create(req)
	if err != nil {
		return err
	}
	res.Vessel = req
	res.Created = true
	return nil
}
