package models

import "time"

type GetUserResponseJSON struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Host           string    `json:"host"`
	ScreenName     string    `json:"screenName"`
	HeaderImageUrl string    `json:"headerImageUrl"`
	IconImageUrl   string    `json:"iconImageUrl"`
	Bio            string    `json:"bio"`
	CreatedAt      time.Time `json:"createdAt"`
}

type GetRemoteUserResponseJSON = GetUserResponseJSON
