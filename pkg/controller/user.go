package controller

import (
	"time"

	"github.com/approvers/qip/pkg/utils/config"

	"github.com/approvers/qip/pkg/controller/models"
	"github.com/approvers/qip/pkg/models/domain"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/usecase"
	"github.com/approvers/qip/pkg/utils/id"
)

// UserController ユーザー関連のAPI
type UserController struct {
	repo        repository.UserRepository
	usecase     usecase.UserUseCase
	idGenerator id.Generator
}

func NewUserController(r repository.UserRepository) *UserController {
	return &UserController{
		repo:        r,
		usecase:     *usecase.NewUserUseCase(r),
		idGenerator: id.NewSnowFlakeIDGenerator(),
	}
}

func (u UserController) CreateUser(q models.CreateUserRequestJSON) (models.CreateUserResponseJSON, error) {
	a := domain.User{
		id:             u.idGenerator.NewID(time.Now()),
		Host:           nil,
		name:           q.Name,
		ScreenName:     q.ScreenName,
		Summary:        "",
		password:       q.Password,
		CreatedAt:      time.Now(),
		UpdatedAt:      nil,
		PrivateKey:     "",
		publicKey:      "",
		WatcherCount:   0,
		WatchingCount:  0,
		PostsCount:     0,
		HeaderImageURL: nil,
		IconImageURL:   nil,
	}

	user, err := u.usecase.Create(a)
	if err != nil {
		return models.CreateUserResponseJSON{}, err
	}

	// 自インスタンスのユーザーである場合はInstanceFQDNで置き換える
	if user.Host == nil {
		f := config.QipConfig.FQDN
		user.Host = &f
	}

	res := models.CreateUserResponseJSON{
		Id:         string(user.id),
		Name:       user.name,
		Host:       *user.Host,
		ScreenName: user.ScreenName,
		CreatedAt:  user.CreatedAt,
	}

	return res, nil
}
