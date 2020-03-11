package model

//用户表
type User struct {
	ID   int    `json:"id"`
	Name string `json:"user_name" db:"UserName"`
	Age  int    `json:"age"`
}

//构造函数
func (u *User) NewUser(id int, username string, age int) *User {
	return &User{
		ID:   id,
		Name: username,
		Age:  age,
	}
}
