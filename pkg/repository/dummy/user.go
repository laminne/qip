package dummy

import (
	"errors"

	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/utils/id"
)

type UserRepository struct {
	data       []domain.User
	followData []domain.Follow
}

func NewUserRepository(data []domain.User, follows []domain.Follow) *UserRepository {
	return &UserRepository{data: data, followData: follows}
}

func (r *UserRepository) FindUsersByName(name string) ([]domain.User, error) {
	res := make([]domain.User, 0)
	for _, v := range r.data {
		if v.GetName() == name {
			res = append(res, v)
		}
	}

	return res, nil
}

func (r *UserRepository) CreateUser(user domain.User) error {
	r.data = append(r.data, user)

	return nil
}

func (r *UserRepository) FindUsersByInstanceID(id id.SnowFlakeID) ([]domain.User, error) {
	res := make([]domain.User, 0)
	for _, v := range r.data {
		if v.GetInstanceID() == id {
			res = append(res, v)
		}
	}

	return res, nil
}

func (r *UserRepository) FindUserByID(id id.SnowFlakeID) (*domain.User, error) {
	for _, v := range r.data {
		if v.GetID() == id {
			return &v, nil
		}
	}

	return nil, errors.New("NotFound")
}

func (r *UserRepository) CreateFollow(f domain.Follow) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) FindUserFollowers(i id.SnowFlakeID) ([]domain.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) UnFollow(from id.SnowFlakeID, target id.SnowFlakeID) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepository) FindUserFollow(i id.SnowFlakeID) ([]domain.Follow, error) {
	res := make([]domain.Follow, 0)
	for _, v := range r.followData {
		if v.GetUserID() == i {
			res = append(res, v)
		}
	}

	return res, nil
}
