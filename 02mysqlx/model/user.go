package model

//用户表
type User struct {
	ID   int64  `json:"id" form:"id" uri:"id"`
	Name string `json:"name" form:"name" uri:"name"`
	Age  int64  `json:"age" form:"age" uri:"age"`
}
