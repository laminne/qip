package entity

import (
	"time"
)

type User struct {
	ID             string  `bun:"id"`
	Host           *string `bun:"host"`
	Name           string  `bun:"name"`
	Password       *string `bun:"password"`
	ScreenName     string  `bun:"screenName"`
	Summary        string  `bun:"summary"`
	PrivateKey     string  `bun:"privateKey"`
	PublicKey      string  `bun:"publicKey"`
	WatcherCount   int     `bun:"watcherCount"`
	WatchingCount  int     `bun:"watchingCount"`
	PostsCount     int     `bun:"postsCount"`
	HeaderImageURL *string `bun:"headerImageUrl"`
	IconImageURL   *string `bun:"iconImageUrl"`

	CreatedAt time.Time  `bun:"createdAt"`
	UpdatedAt *time.Time `bun:"updatedAt"`
}
