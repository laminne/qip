package models

import "time"

type CreateUserRequestJSON struct {
	Name       string `json:"name"`
	ScreenName string `json:"screenName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type CreateUserResponseJSON struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Host       string    `json:"host"`
	ScreenName string    `json:"screenName"`
	CreatedAt  time.Time `json:"createdAt"`
}
