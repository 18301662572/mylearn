package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //驱动，连接数据库必备
)

//连接mysql

//全局变量
var db *sql.DB

//InitDB 初始化数据库
func InitDB() (err error) {
	//DSN:Data Source Name
	dsn := "root:123456@tcp(192.168.42.133:3306)/mytest"
	//这里不会校验账号密码是否正确
	//注意！！！这里不要使用:=,我们是给全局变量赋值
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("connect to mysql failed")
		return err
	}
	fmt.Println("connect to mysql success")
	//defer db.Close()
	//尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		fmt.Println("ping mysql failed")
		return err
	}
	return nil
}
