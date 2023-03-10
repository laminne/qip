package user

import (
	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/errorType"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type IFindUserService interface {
	FindByID(id id.SnowFlakeID) (*UserData, error)
	FindByInstanceID(instanceID id.SnowFlakeID) ([]UserData, error)
	FindByName(name string) ([]UserData, error)
}

type FindUserService struct {
	userRepository repository.UserRepository
}

func NewFindUserService(repo repository.UserRepository) *FindUserService {
	return &FindUserService{userRepository: repo}
}

func (f *FindUserService) convert(u []domain.User) []UserData {
	d := make([]UserData, len(u))

	for i, v := range u {
		d[i] = *NewUserData(v)
	}
	return d
}

func (f *FindUserService) FindByID(id id.SnowFlakeID) (*UserData, error) {
	u, err := f.userRepository.FindUserByID(id)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindUserService", "user not found")
	}

	return NewUserData(*u), nil
}

func (f *FindUserService) FindByInstanceID(instanceID id.SnowFlakeID) ([]UserData, error) {
	u, err := f.userRepository.FindUsersByInstanceID(instanceID)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindUserService", "no such instance")
	}

	return f.convert(u), nil
}

func (f *FindUserService) FindByName(name string) ([]UserData, error) {
	u, err := f.userRepository.FindUsersByName(name)
	if err != nil {
		return nil, errorType.NewErrNotFound("FindUserService", "user not found")
	}

	return f.convert(u), nil
}
