package entity

import "time"

type Post struct {
	ID         string    `bun:"id"`
	Body       string    `bun:"body"`
	Type       string    `bun:"type"`
	MergeCount int       `bun:"mergeCount"`
	Visibility string    `bun:"visibility"`
	CreatedAt  time.Time `bun:"createdAt"`

	UserID string `bun:"userID"`
}
