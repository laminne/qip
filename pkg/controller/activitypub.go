package controller

import (
	"strconv"

	"github.com/approvers/qip/pkg/utils/id"

	"github.com/approvers/qip/pkg/activitypub"
	"github.com/approvers/qip/pkg/activitypub/types"
	"github.com/approvers/qip/pkg/models/domain"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/usecase"
)

type ActivityPubController struct {
	repo    repository.UserRepository
	usecase usecase.UserUseCase
}

func NewActivityPubController(r repository.UserRepository) *ActivityPubController {
	return &ActivityPubController{
		repo:    r,
		usecase: *usecase.NewUserUseCase(r),
	}
}

func (c ActivityPubController) GetUser(uid string) *types.PersonResponseJSONLD {
	// snowflakeかUsernameか判別
	_, err := strconv.Atoi(uid)
	var user *domain.User
	if err != nil {
		// UserNameのとき
		user, err = c.usecase.FindLocalByUserName(uid)
		if err != nil {
			return nil
		}
	} else {
		// SnowflakeIDのとき
		user, err = c.usecase.FindByID(id.SnowFlakeID(uid))
		if err != nil {
			return nil
		}
	}

	if user == nil {
		return nil
	}

	n := ""
	if user.IconImageURL == nil {
		user.IconImageURL = &n
	}
	if user.HeaderImageURL == nil {
		user.HeaderImageURL = &n
	}

	arg := types.PersonResponseArgs{
		ID:             string(user.ID),
		UserName:       user.Name,
		UserScreenName: user.ScreenName,
		Summary:        user.Summary,
		Icon: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			Url: *user.IconImageURL,
		},
		Image: struct {
			Url       string
			Sensitive bool
			Name      interface{}
		}{
			Url: *user.HeaderImageURL,
		},
		Tag:                       nil,
		ManuallyApprovesFollowers: false,
		PublicKey:                 user.PublicKey,
	}

	res := activitypub.Person(arg)
	return &res
}
