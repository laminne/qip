package models

type LoginRequestJSON struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponseJSON struct {
	Token string `json:"token"`
}
