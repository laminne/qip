package models

type FindUserFollowersResponseJSON struct {
	Followers []string `json:"followers"`
}

type FindUserFollowResponseJSON struct {
	Follows []string `json:"follows"`
}
