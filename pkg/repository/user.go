package repository

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type UserRepository interface {
	FindUsersByName(name string) ([]domain.User, error)
	FindUserByID(id id.SnowFlakeID) (*domain.User, error)
	FindUsersByInstanceID(id id.SnowFlakeID) ([]domain.User, error)
	CreateUser(u domain.User) error
}
