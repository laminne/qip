package domain

import (
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

// Reaction リアクション
type Reaction struct {
	ReactedUserID id.SnowFlakeID // リアクションしたユーザー
	PostID        id.SnowFlakeID // リアクション先の投稿
	CreatedAt     time.Time
}
