package usecase

import (
	"fmt"

	"github.com/laminne/notepod/pkg/models/domain"
	"github.com/laminne/notepod/pkg/repository"
	"github.com/laminne/notepod/pkg/utils/id"
)

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	fmt.Println(repo)
	return &UserUseCase{repo: repo}
}

func (c UserUseCase) FindByID(id id.SnowFlakeID) (*domain.User, error) {
	u, err := c.repo.FindByID(id)
	if err != nil || u == nil {
		return nil, err
	}
	return u, nil
}

func (c UserUseCase) FindByUserName(name string) ([]domain.User, error) {
	u, err := c.repo.FindByUserName(name)
	if err != nil {
		return []domain.User{}, err
	}
	return u, nil
}

func (c UserUseCase) FindLocalByUserName(name string) (*domain.User, error) {
	u, err := c.repo.FindLocalByUserName(name)
	if err != nil {
		fmt.Println(u, err)
		return nil, err
	}

	return &u, nil
}

func (c UserUseCase) Create(u domain.User) (domain.User, error) {
	r, err := c.repo.Create(u)
	if err != nil {
		return domain.User{}, err
	}
	return r, nil
}
