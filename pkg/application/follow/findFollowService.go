package follow

import (
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/id"
)

type IFindFollowService interface {
	FindUserFollow(i id.SnowFlakeID) ([]FollowData, error)
	FindUserFollowers(i id.SnowFlakeID) ([]FollowData, error)
}

type FindFollowService struct {
	repository repository.FollowRepository
}

func (f *FindFollowService) FindUserFollow(i id.SnowFlakeID) ([]FollowData, error) {
	d, err := f.repository.FindUserFollow(i)
	if err != nil {
		return nil, err
	}

	res := make([]FollowData, len(d))
	for i, v := range d {
		res[i] = *NewFollowData(v)
	}

	return res, nil
}

func (f *FindFollowService) FindUserFollowers(i id.SnowFlakeID) ([]FollowData, error) {
	d, err := f.repository.FindUserFollowers(i)
	if err != nil {
		return nil, err
	}

	res := make([]FollowData, len(d))
	for i, v := range d {
		res[i] = *NewFollowData(v)
	}

	return res, nil
}

func NewFindFollowService(repo repository.FollowRepository) *FindFollowService {
	return &FindFollowService{repository: repo}
}
