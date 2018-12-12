package model

type User struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
