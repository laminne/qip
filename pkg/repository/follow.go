package repository

import (
	"github.com/laminne/qip/pkg/domain"
	"github.com/laminne/qip/pkg/utils/id"
)

type FollowRepository interface {
	Create(f domain.Follow) error
	FindUserFollowers(i id.SnowFlakeID) ([]domain.Follow, error)
	FindUserFollow(i id.SnowFlakeID) ([]domain.Follow, error)
	UnFollow(from id.SnowFlakeID, target id.SnowFlakeID) error
}
