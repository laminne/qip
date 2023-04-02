package models

import "time"

type GetPostResponseJSON struct {
	ID         string                    `json:"id"`
	Body       string                    `json:"body"`
	AuthorID   string                    `json:"authorID"`
	Visibility string                    `json:"visibility"`
	User       GetPostResponseAuthorData `json:"user"`
	CreatedAt  time.Time                 `json:"createdAt"`
}

type GetPostResponseAuthorData struct {
	Name       string `json:"name"`
	Host       string `json:"host"`
	ScreenName string `json:"screenName"`
	IconURL    string `json:"iconURL"`
}
