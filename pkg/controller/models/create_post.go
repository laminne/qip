package models

import (
	"time"
)

type CreatePostRequestJSON struct {
	Body       string `json:"body"`
	Visibility int    `json:"visibility"`
}

type CreatePostResponseJSON struct {
	ID         string    `json:"id"`
	Body       string    `json:"body"`
	AuthorID   string    `json:"authorID"`
	Visibility string    `json:"visibility"`
	CreatedAt  time.Time `json:"createdAt"`
}
