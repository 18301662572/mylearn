package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"jmoiron/sqlx"
)

//db:数据库连接

//声明数据库实例
var (
	db *sqlx.DB
)

//初始化
func Init(mysqlDNS string) {
	//获取连接
	db, err := sqlx.Connect("mysql", mysqlDNS)
	if err != nil {
		fmt.Printf("connect sqlx failed,err:%v\n", err)
	}
	//设置闲置的连接数
	db.SetMaxIdleConns(1)
	//设置最大打开的连接数
	db.SetMaxOpenConns(3)
}
