package repository

import "github.com/approvers/qip/pkg/domain"

type UserRepository interface {
	FindUsersByName(name string) ([]domain.User, error)
	CreateUser(u domain.User) error
}
