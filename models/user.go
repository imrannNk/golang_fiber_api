package models

type User struct {
	UserId      string `json:"UserId"`
	UserName    string `json:"UserName"`
	UserAge     int    `json:"UserAge"`
	UserAddress string `json:"UserAddress"`
}
