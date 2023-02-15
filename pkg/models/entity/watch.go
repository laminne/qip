package entity

import "time"

type Watch struct {
	UserID    string    `bun:"userID"`
	TargetID  string    `bun:"targetID"`
	CreatedAt time.Time `bun:"createdAt"`
}
