package domain

import (
	"time"

	"github.com/approvers/qip/pkg/utils/id"
)

// Watch ウォッチ(フォロー)
type Watch struct {
	UserID    id.SnowFlakeID
	TargetID  id.SnowFlakeID
	CreatedAt time.Time
}
