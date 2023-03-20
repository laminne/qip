package controller

import (
	"fmt"

	"github.com/approvers/qip/pkg/application/file"
	"github.com/approvers/qip/pkg/application/instance"
	"github.com/approvers/qip/pkg/utils"

	"github.com/approvers/qip/pkg/controller/models"

	"github.com/approvers/qip/pkg/application/user"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

// UserController ユーザー関連のAPI
type UserController struct {
	repo                repository.UserRepository
	findUserService     user.FindUserService
	findFileService     file.FindFileService
	findInstanceService instance.FindInstanceService
}

func NewUserController(r repository.UserRepository, f repository.FileRepository, i repository.InstanceRepository) *UserController {
	return &UserController{
		repo:                r,
		findUserService:     *user.NewFindUserService(r),
		findFileService:     *file.NewFindFileService(f),
		findInstanceService: *instance.NewFindInstanceService(i),
	}
}

func (u *UserController) FindUserByID(id id.SnowFlakeID) (models.GetUserResponseJSON, error) {
	user, err := u.findUserService.FindByID(id)
	if err != nil {
		return models.GetUserResponseJSON{}, err
	}
	var (
		headerURL string
		iconURL   string
	)
	if user.HeaderImageID() != nil {
		data, err := u.findFileService.FindByID(*user.HeaderImageID())
		if err != nil {
			return models.GetUserResponseJSON{}, err
		}
		headerURL = data.FileURL()
		fmt.Println("Set", headerURL)
	}
	if user.IconImageID() != nil {
		data, err := u.findFileService.FindByID(*user.IconImageID())
		if err != nil {
			return models.GetUserResponseJSON{}, err
		}
		iconURL = data.FileURL()
		fmt.Println("Set", iconURL)
	}
	i, err := u.findInstanceService.FindByID(user.InstanceID())
	if err != nil {
		return models.GetUserResponseJSON{}, err
	}

	n := utils.NilFiller[string](user.Bio())

	return models.GetUserResponseJSON{
		Id:             string(user.Id()),
		Name:           user.Name(), // ToDo: Host情報も返却する
		Host:           i.Host(),
		ScreenName:     user.DisplayName(),
		HeaderImageUrl: headerURL,
		IconImageUrl:   iconURL,
		Bio:            n[0],
		CreatedAt:      user.CreatedAt(),
	}, nil
}
