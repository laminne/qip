package entity

import (
	"time"
)

type User struct {
	ID             string     `bun:"id"`
	Host           *string    `bun:"host"`
	Name           string     `bun:"name"`
	ScreenName     string     `bun:"screenName"`
	Summary        string     `bun:"summary"`
	CreatedAt      time.Time  `bun:"createdAt"`
	UpdatedAt      *time.Time `bun:"updatedAt"`
	PrivateKey     string     `bun:"privateKey"`
	PublicKey      string     `bun:"publicKey"`
	FollowerCount  int        `bun:"followerCount"`
	FollowingCount int        `bun:"followingCount"`
	NoteCount      int        `bun:"noteCount"`
	HeaderImageURL *string    `bun:"headerImageUrl"`
	IconImageURL   *string    `bun:"iconImageUrl"`
}
