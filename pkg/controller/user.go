package controller

import (
	"errors"

	"github.com/laminne/qip/pkg/errorType"

	"github.com/laminne/qip/pkg/activitypub"

	"github.com/laminne/qip/pkg/application/file"
	"github.com/laminne/qip/pkg/application/instance"
	"github.com/laminne/qip/pkg/utils"

	"github.com/laminne/qip/pkg/controller/models"

	"github.com/laminne/qip/pkg/application/user"
	"github.com/laminne/qip/pkg/repository"
	"github.com/laminne/qip/pkg/utils/id"
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
	}
	if user.IconImageID() != nil {
		data, err := u.findFileService.FindByID(*user.IconImageID())
		if err != nil {
			return models.GetUserResponseJSON{}, err
		}
		iconURL = data.FileURL()
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

func (u *UserController) FindUserByAcct(acct string) (models.GetUserResponseJSON, error) {
	// acctParserでacctを分解
	a := activitypub.AcctParser(acct)
	// ユーザー名で検索
	us, err := u.findUserService.FindByName(a.UserName)
	if err != nil {
		return models.GetUserResponseJSON{}, err
	}
	if a.Host == nil {
		return models.GetUserResponseJSON{}, errors.New("failed to parse acct")
	}
	// acctのホストをInstanceから検索
	in, err := u.findInstanceService.FindByHost(*a.Host)
	if err != nil {
		return models.GetUserResponseJSON{}, err
	}
	// ユーザーのInstanceIDでマッチするユーザーを検索
	for _, v := range us {
		if v.InstanceID() == in.Id() && in.Host() == *a.Host {
			return u.FindUserByID(v.Id())
		}
	}

	return models.GetUserResponseJSON{}, errorType.NewErrNotFound("UserController", "user not found")
}
