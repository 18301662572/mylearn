package models

import (
	"time"
)

type Book struct {
	BookCategoryId int64     `json:"book_category_id" xorm:"not null comment('书籍类型ID') index BIGINT(20)"`
	BookName       string    `json:"book_name" xorm:"not null comment('书籍名称') index VARCHAR(64)"`
	BookNo         int       `json:"book_no" xorm:"not null comment('书籍排序') INT(11)"`
	BookState      int       `json:"book_state" xorm:"not null default b'0' comment('书籍状态（0：存在 1：删除）') BIT(1)"`
	Content        string    `json:"content" xorm:"comment('书籍内容') index LONGTEXT"`
	CreateTime     time.Time `json:"create_time" xorm:"not null comment('创建时间') TIMESTAMP"`
	Id             int64     `json:"id" xorm:"pk autoincr comment('主键ID') BIGINT(20)"`
	UpdateTime     time.Time `json:"update_time" xorm:"comment('修改时间') TIMESTAMP"`
}
