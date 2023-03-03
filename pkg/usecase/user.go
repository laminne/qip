package usecase

import (
	"fmt"

	"github.com/approvers/qip/pkg/utils/password/argon2"

	"github.com/approvers/qip/pkg/utils/password"

	"github.com/approvers/qip/pkg/models/domain"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type UserUseCase struct {
	repo    repository.UserRepository
	encoder password.Encoder
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	fmt.Println(repo)
	return &UserUseCase{repo: repo, encoder: argon2.NewArgon2PasswordEncoder()}
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
	encodedPassword, err := c.encoder.EncodePassword(u.password)
	if err != nil {
		return domain.User{}, err
	}

	u.password = string(encodedPassword)
	fmt.Println(u.password)

	r, err := c.repo.Create(u)
	if err != nil {
		return domain.User{}, err
	}
	return r, nil
}
