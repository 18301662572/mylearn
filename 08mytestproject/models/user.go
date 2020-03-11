package models

import (
	"time"
)

type User struct {
	CreateTime time.Time `json:"create_time" xorm:"not null comment('创建时间') TIMESTAMP"`
	Id         int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(20)"`
	NickName   string    `json:"nick_name" xorm:"comment('昵称') VARCHAR(64)"`
	Password   string    `json:"password" xorm:"not null comment('密码') VARCHAR(64)"`
	UpdateTime time.Time `json:"update_time" xorm:"comment('修改时间') TIMESTAMP"`
	UserId     int64     `json:"user_id" xorm:"not null comment('用户ID') unique BIGINT(20)"`
	UserName   string    `json:"user_name" xorm:"not null comment('登录名称') unique VARCHAR(64)"`
	UserState  int       `json:"user_state" xorm:"not null default b'0' comment('用户状态（0：存在 1：删除）') BIT(1)"`
}
