package main

import (
	"github.com/jinzhu/gorm"
	pb "github.com/lty5240/consignment/shippy-service-user/proto/user"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmailAndPassword(user *pb.User) (*pb.User, error)
	GetByEmail(email string) (*pb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repository *UserRepository) GetAll() ([]*pb.User, error) {
	var users []*pb.User
	if err := repository.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepository) Get(id string) (*pb.User, error) {
	user := &pb.User{}
	user.Id = id
	if err := repository.db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) GetByEmailAndPassword(user *pb.User) (*pb.User, error) {
	if err := repository.db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repository *UserRepository) Create(user *pb.User) error {
	if err := repository.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) GetByEmail(email string) (*pb.User, error) {
	user := &pb.User{
		Email: email,
	}
	if err := repository.db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
