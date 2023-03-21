package follow

import (
	"time"

	"github.com/approvers/qip/pkg/domain"
	"github.com/approvers/qip/pkg/utils/id"
)

type FollowData struct {
	userID    id.SnowFlakeID
	targetID  id.SnowFlakeID
	createdAt time.Time
}

func NewFollowData(f domain.Follow) *FollowData {
	return &FollowData{
		f.GetUserID(),
		f.GetTargetID(),
		f.GetCreatedAt(),
	}
}

func (f FollowData) UserID() id.SnowFlakeID {
	return f.userID
}

func (f FollowData) TargetID() id.SnowFlakeID {
	return f.targetID
}

func (f FollowData) CreatedAt() time.Time {
	return f.createdAt
}
