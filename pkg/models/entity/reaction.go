package entity

import "time"

type Reaction struct {
	ReactedUserID string    `bun:"reactedUserID"`
	PostID        string    `bun:"postID"`
	CreatedAt     time.Time `bun:"createdAt"`
}
