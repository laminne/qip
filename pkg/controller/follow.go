package controller

import (
	"github.com/approvers/qip/pkg/application/follow"
	"github.com/approvers/qip/pkg/controller/models"
	"github.com/approvers/qip/pkg/domain/service"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type FollowController struct {
	repo                repository.FollowRepository
	createFollowService follow.ICreateFollowService
	findFollowService   follow.IFindFollowService
}

func NewFollowController(repo repository.FollowRepository) *FollowController {
	return &FollowController{
		repo:                repo,
		createFollowService: follow.NewCreateFollowService(*service.NewFollowService(repo), repo),
		findFollowService:   follow.NewFindFollowService(repo),
	}
}

// FindUserFollower 特定ユーザーのフォロワーを取得
func (f *FollowController) FindUserFollower(i id.SnowFlakeID) (models.FindUserFollowersResponseJSON, error) {
	r, err := f.findFollowService.FindUserFollowers(i)
	if err != nil {
		return models.FindUserFollowersResponseJSON{}, err
	}
	res := make([]string, len(r))
	for j, k := range r {
		res[j] = string(k.UserID())
	}

	return models.FindUserFollowersResponseJSON{Followers: res}, nil
}

// FindUserFollow 特定ユーザーがフォローしているユーザーを取得
func (f *FollowController) FindUserFollow(uID id.SnowFlakeID) (models.FindUserFollowResponseJSON, error) {
	res, err := f.findFollowService.FindUserFollow(uID)
	if err != nil {
		return models.FindUserFollowResponseJSON{}, err
	}
	r := make([]string, len(res))
	for i, v := range res {
		r[i] = string(v.TargetID())
	}

	return models.FindUserFollowResponseJSON{Follows: r}, nil
}

func (f *FollowController) Create(userID id.SnowFlakeID, target id.SnowFlakeID) error {
	c := follow.CreateFollowCommand{
		UserID:   userID,
		TargetID: target,
	}

	_, err := f.createFollowService.Handle(c)
	if err != nil {
		return err
	}
	return nil
}
