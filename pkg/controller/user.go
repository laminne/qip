package controller

import (
	"errors"

	"github.com/approvers/qip/pkg/utils"

	"github.com/approvers/qip/pkg/controller/models"

	"github.com/approvers/qip/pkg/application/user"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

// UserController ユーザー関連のAPI
type UserController struct {
	repo            repository.UserRepository
	findUserService user.FindUserService
}

func NewUserController(r repository.UserRepository) *UserController {
	return &UserController{
		repo:            r,
		findUserService: *user.NewFindUserService(r),
	}
}

func (u *UserController) FindUserByID(id id.SnowFlakeID) (models.GetUserResponseJSON, error) {
	user, err := u.findUserService.FindByID(id)
	if err != nil {
		return models.GetUserResponseJSON{}, errors.New("")
	}

	n := utils.NilFiller[string]((*string)(user.HeaderImageID()), (*string)(user.IconImageID()), user.Bio())

	return models.GetUserResponseJSON{
		Id:             string(user.Id()),
		Name:           user.Name(), // ToDo: Host情報も返却する
		ScreenName:     user.DisplayName(),
		HeaderImageUrl: n[0],
		IconImageUrl:   n[1],
		Bio:            n[2],
		CreatedAt:      user.CreatedAt(),
	}, nil
}
