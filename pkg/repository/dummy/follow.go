package dummy

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/utils/id"
)

type FollowRepository struct {
	data []domain.Follow
}

func (r *FollowRepository) Create(f domain.Follow) error {
	//TODO implement me
	panic("implement me")
}

func (r *FollowRepository) FindUserFollowers(i id.SnowFlakeID) ([]domain.Follow, error) {
	//TODO implement me
	panic("implement me")
}

func (r *FollowRepository) UnFollow(from id.SnowFlakeID, target id.SnowFlakeID) error {
	//TODO implement me
	panic("implement me")
}

func NewFollowRepository(data []domain.Follow) *FollowRepository {
	return &FollowRepository{data: data}
}

func (r *FollowRepository) FindUserFollow(i id.SnowFlakeID) ([]domain.Follow, error) {
	res := make([]domain.Follow, 0)
	for _, v := range r.data {
		if v.GetUserID() == i {
			res = append(res, v)
		}
	}

	return res, nil
}
