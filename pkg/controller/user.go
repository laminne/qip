package controller

import (
	"time"

	"github.com/laminne/notepod/pkg/activitypub"

	"github.com/laminne/notepod/pkg/controller/models"
	"github.com/laminne/notepod/pkg/models/domain"
	"github.com/laminne/notepod/pkg/repository"
	"github.com/laminne/notepod/pkg/usecase"
	"github.com/laminne/notepod/pkg/utils/id"
)

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
		ID:             u.idGenerator.NewID(time.Now()),
		Host:           nil,
		Name:           q.Name,
		ScreenName:     q.ScreenName,
		Summary:        "",
		Password:       q.Password,
		CreatedAt:      time.Now(),
		UpdatedAt:      nil,
		PrivateKey:     "",
		PublicKey:      "",
		FollowerCount:  0,
		FollowingCount: 0,
		NoteCount:      0,
		HeaderImageURL: nil,
		IconImageURL:   nil,
	}

	user, err := u.usecase.Create(a)
	if err != nil {
		return models.CreateUserResponseJSON{}, err
	}

	// 自インスタンスのユーザーである場合はInstanceFQDNで置き換える
	if user.Host == nil {
		f := activitypub.InstanceFQDN
		user.Host = &f
	}

	res := models.CreateUserResponseJSON{
		Id:         string(user.ID),
		Name:       user.Name,
		Host:       *user.Host,
		ScreenName: user.ScreenName,
		CreatedAt:  user.CreatedAt,
	}

	return res, nil
}
