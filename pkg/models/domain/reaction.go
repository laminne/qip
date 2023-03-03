package domain

import (
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

// Reaction リアクション
type Reaction struct {
	reactedUserID id.SnowFlakeID // リアクションしたユーザー
	targetID      id.SnowFlakeID // リアクション先の投稿
	createdAt     time.Time
}

func NewReaction(reacted id.SnowFlakeID, target id.SnowFlakeID, now time.Time) (*Reaction, error) {
	return &Reaction{
		reactedUserID: reacted,
		targetID:      target,
		createdAt:     now,
	}, nil
}

func (r *Reaction) GetReactedUserID() id.SnowFlakeID {
	return r.reactedUserID
}

func (r *Reaction) GetTargetID() id.SnowFlakeID {
	return r.targetID
}

func (r *Reaction) GetCreatedAt() time.Time {
	return r.createdAt
}
