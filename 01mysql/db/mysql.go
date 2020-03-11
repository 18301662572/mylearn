package db

import (
	"code.oldbody.com/studygolang/mylearn/01mysql/model"
	"fmt"
)

//操作mysql

//添加用户
func AddUser(username string, age int) (theid int64, err error) {
	sqlStr := "insert into user(username,age) value(?,?)"
	res, err := db.Exec(sqlStr, username, age)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return 0, err
	}
	theid, err = res.LastInsertId() //新插入数据的ID
	if err != nil {
		fmt.Printf("get LastInsertId failed,err:%v\n", err)
		return 0, err
	}
	fmt.Printf("insert success,the id is %d\n", theid)
	return theid, nil
}

//修改用户
func UpdataUser(age int, id int64) error {
	sqlStr := "update user set age=? where id=?"
	res, err := db.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return err
	}
	rows, err := res.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return err
	}
	fmt.Printf("update success,affected rows is %d\n", rows)
	return nil
}

func DeleteUser(id int) error {
	sqlStr := "delete from user where id=?"
	res, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed,err:%v\n", err)
		return err
	}
	rows, err := res.RowsAffected() //操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed,err:%v\n", err)
		return err
	}
	fmt.Printf("delete success,affected rows is %d\n", rows)
	return nil
}

//查询单行用户
func QueryUser(id int64) (*model.User, error) {
	var u model.User
	sqlStr := "select id,username,age from user where id=?"
	//非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		fmt.Printf("scan failed,err:%v\n", err)
		return nil, err
	}
	fmt.Printf("id:%d username:%s age%d \n", u.ID, u.Name, u.Age)
	return &u, nil
}

//查询多行用户
func QueryUsers() error {
	sqlStr := "select id,username,age from user"
	//非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return err
	}
	//非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	//读取数据
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return err
		}
		fmt.Printf("id:%d username:%s age%d \n", u.ID, u.Name, u.Age)
	}
	return nil
}
