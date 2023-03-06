package dummy

import (
	"errors"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type UserRepository struct {
	data []domain.User
}

func NewUserRepository(data []domain.User) *UserRepository {
	return &UserRepository{data: data}
}

func (u *UserRepository) FindUsersByName(name string) ([]domain.User, error) {
	res := make([]domain.User, 0)
	for _, v := range u.data {
		if v.GetName() == name {
			res = append(res, v)
		}
	}

	return res, nil
}

func (u *UserRepository) CreateUser(user domain.User) error {
	u.data = append(u.data, user)

	return nil
}

func (u *UserRepository) FindUsersByInstanceID(id id.SnowFlakeID) ([]domain.User, error) {
	res := make([]domain.User, 0)
	for _, v := range u.data {
		if v.GetInstanceID() == id {
			res = append(res, v)
		}
	}

	return res, nil
}

func (u *UserRepository) FindUserByID(id id.SnowFlakeID) (*domain.User, error) {
	for _, v := range u.data {
		if v.GetID() == id {
			return &v, nil
		}
	}

	return nil, errors.New("NotFound")
}
