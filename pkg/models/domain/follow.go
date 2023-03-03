package domain

import (
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

// Follow フォロー
type Follow struct {
	userID    id.SnowFlakeID
	targetID  id.SnowFlakeID
	createdAt time.Time
}

func NewFollow(userID id.SnowFlakeID, targetID id.SnowFlakeID, now time.Time) (*Follow, error) {
	return &Follow{
		userID:    userID,
		targetID:  targetID,
		createdAt: now,
	}, nil
}

func (f *Follow) GetUserID() id.SnowFlakeID {
	return f.userID
}

func (f *Follow) GetTargetID() id.SnowFlakeID {
	return f.targetID
}

func (f *Follow) GetCreatedAt() time.Time {
	return f.createdAt
}
