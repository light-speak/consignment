package user

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

func (*User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()
	return scope.SetColumn("Id", id.String())
}
