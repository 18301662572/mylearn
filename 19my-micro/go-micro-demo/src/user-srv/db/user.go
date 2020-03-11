package db

import (
	"code.oldbody.com/studygolang/mylearn/19my-micro/go-micro-demo/src/user-srv/entity"
	"database/sql"
)

//DAL :数据库访问层

//user 表的查询
func SelectUserById(id int32) (*entity.User, error) {
	//创建user对象，用于返回数据
	user := entity.User{}
	//执行查询
	err := db.Get(&user, "SELECT id,name,address,phone FROM user WHERE id=?", id)
	if err == sql.ErrNoRows { //查询行会出现的异常
		return nil, err
	}
	return &user, err
}

//user 数据的增加
func InsertUser(user *entity.User) (int64, error) {
	rep, err := db.Exec("INSERT INTO user(name,address,phone) VALUE(?,?,?)", user.Name, user.Address, user.Phone)
	if err != nil {
		return 0, err
	}
	return rep.LastInsertId()
}

//修改
func ModifyUser(user *entity.User) error {
	_, err := db.Exec("UPDATE user SET name=?,address=?,phone=? WHERE id=?",
		user.Name, user.Address, user.Phone, user.Id)
	if err != nil {
		return err
	}
	return nil
}

//删除
func DeleteUser(id int32) error {
	_, err := db.Exec("DELETE FROM user WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}
