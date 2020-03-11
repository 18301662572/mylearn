package main

import (
	"code.oldbody.com/studygolang/mylearn/01mysql/db"
	"fmt"
)

//mysql

func main() {
	err := db.InitDB() //调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}

	//添加用户
	theid, err := db.AddUser("张三", 20)
	if err != nil {
		fmt.Printf("AddUser failed,err:%v\n", err)
		return
	}
	//修改
	err = db.UpdataUser(29, theid)
	if err != nil {
		fmt.Printf("UpdataUser failed,err:%v\n", err)
		return
	}
	//删除
	//err=db.DeleteUser(1)
	//if err!=nil{
	//	fmt.Printf("DeleteUser failed,err:%v\n",err)
	//	return
	//}
	//查询单行
	u, err := db.QueryUser(theid)
	if err != nil {
		fmt.Printf("QueryUser failed,err:%v\n", err)
		return
	}
	fmt.Printf("id:%d username:%s age%d \n", u.ID, u.Name, u.Age)
	//查询多行
	err = db.QueryUsers()
	if err != nil {
		fmt.Printf("QueryUsers failed,err:%v\n", err)
		return
	}
	fmt.Println("QueryUsers success")
}
