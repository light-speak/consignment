package main

import (
	"context"
	"errors"
	pb "github.com/lty5240/consignment/shippy-service-user/proto/user"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type service struct {
	repository   Repository
	tokenService Authable
}

func (s service) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	if err := s.repository.Create(req); err != nil {
		return err
	}
	res.User = req
	return nil
}

func (s service) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	user, err := s.repository.Get(req.Id)
	if err != nil {
		return err
	}
	res.User = user
	return nil
}

func (s service) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	users, err := s.repository.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return nil
}

func (s service) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	log.Println("Logging in with : ", req.Email, req.Password)
	user, err := s.repository.GetByEmail(req.Email)
	log.Println(user)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil
	}
	token, err := s.tokenService.Encode(user)
	if err != nil {
		return err
	}
	res.Token = token
	return nil
}

func (s service) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := s.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	log.Println(claims)

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}
	res.Valid = true
	return nil
}
