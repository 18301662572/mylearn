package model

type User struct {
	UserName string `form:"username"`
	PassWord string `form:"pwd"`
}
