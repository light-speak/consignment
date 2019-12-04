package main

import (
	"context"
	"github.com/lty5240/consignment/shippy-service-consignment/proto/consignment"
	"github.com/lty5240/consignment/shippy-service-vessel/proto/vessel"
	"log"
)

type handler struct {
	repository
	vesselClient vessel.VesselServiceClient
}

func (s *handler) CreateConsignment(ctx context.Context, req *consignment.Consignment, res *consignment.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vessel.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	log.Printf("Found vessel : %s \n", vesselResponse.Vessel.Name)
	if err != nil {
		return err
	}

	req.VesselId = vesselResponse.Vessel.Id
	if err = s.repository.Create(req); err != nil {
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}


func (s *handler) GetConsignments(ctx context.Context, req *consignment.GetRequest, res *consignment.Response) error {
	consignments, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	res.Consignments = consignments
	return nil
}

