package user

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type IFindUserService interface {
	FindByID(id id.SnowFlakeID) (*domain.User, error)
	FindByInstanceID(instanceID id.SnowFlakeID) ([]domain.User, error)
	FindByName(name string) ([]domain.User, error)
}

type FindUserService struct {
	userRepository repository.UserRepository
}

func NewFindUserService(repo repository.UserRepository) *FindUserService {
	return &FindUserService{userRepository: repo}
}

func (f *FindUserService) FindByID(id id.SnowFlakeID) (*domain.User, error) {
	u, err := f.userRepository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (f *FindUserService) FindByInstanceID(instanceID id.SnowFlakeID) ([]domain.User, error) {
	u, err := f.userRepository.FindUsersByInstanceID(instanceID)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (f *FindUserService) FindByName(name string) ([]domain.User, error) {
	u, err := f.userRepository.FindUsersByName(name)
	if err != nil {
		return nil, err
	}

	return u, nil
}
