package models

type User struct {
	Uuid     int    `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
}
