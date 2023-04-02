package repository

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/utils/id"
)

type UserRepository interface {
	FindUsersByName(name string) ([]domain.User, error)
	FindUserByID(id id.SnowFlakeID) (*domain.User, error)
	FindUsersByInstanceID(id id.SnowFlakeID) ([]domain.User, error)
	CreateUser(u domain.User) error
}
