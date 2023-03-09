package models

import "time"

type GetPostResponseJSON struct {
	ID         string    `json:"id"`
	Body       string    `json:"body"`
	AuthorID   string    `json:"authorID"`
	Visibility string    `json:"visibility"`
	CreatedAt  time.Time `json:"createdAt"`
}
