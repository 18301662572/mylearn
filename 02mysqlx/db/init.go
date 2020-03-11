package db

import (
	"code.oldbody.com/studygolang/mylearn/02mysqlx/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//sqlx

var db *sqlx.DB

//InitDB 初始化数据库
func InitDB() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect sqlx failed,err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return nil
}

//添加用户
func AddUser(age int64, name string) (theid int64, err error) {
	sqlStr := "insert into user(name,age) value(?,?)"
	res, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert user failed,err:%v\n", err)
		return 0, err
	}
	theid, err = res.LastInsertId()
	if err != nil {
		fmt.Printf("get LastInsertId failed,err:%v\n", err)
		return theid, err
	}
	return theid, nil
}

//删除用户
func DeleteUser(id int64) (int64, error) {
	sqlStr := "delete from user where id=?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete user failed,err:%v\n", err)
		return id, err
	}
	return id, nil
}

//修改用户
func UpdateUser(name string, age int64, id int64) (int64, error) {
	sqlStr := "update user set age=?,name=? where id=?"
	_, err := db.Exec(sqlStr, age, name, id)
	if err != nil {
		fmt.Printf("update user failed,err:%v\n", err)
		return id, err
	}
	return id, nil
}

//查询单个用户
func QueryUser(id int64) (*model.User, error) {
	var u model.User
	sqlStr := "select id,name,age from user where id=?"
	//非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Printf("select user demo failed,err:%v\n", err)
		return nil, err
	}
	return &u, nil
}

//查询多个用户
func QueryUserList() (*[]model.User, error) {
	var u []model.User
	sqlStr := "select id,name,age from user"
	//非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.Select(&u, sqlStr)
	if err != nil {
		fmt.Printf("select userlist failed,err:%v\n", err)
		return nil, err
	}
	return &u, nil
}
