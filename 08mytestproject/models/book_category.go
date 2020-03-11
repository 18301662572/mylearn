package models

import (
	"time"
)

type BookCategory struct {
	BookCategoryName  string    `json:"book_category_name" xorm:"not null comment('书籍分类名称') index VARCHAR(64)"`
	BookCategoryNo    int       `json:"book_category_no" xorm:"not null comment('书籍分类排序') INT(11)"`
	BookCategoryState int       `json:"book_category_state" xorm:"not null default b'0' comment('书籍分类状态（0：存在 1：删除）') BIT(1)"`
	Createtime        time.Time `json:"createtime" xorm:"not null comment('创建时间') TIMESTAMP"`
	Id                int64     `json:"id" xorm:"pk autoincr comment('书籍分类ID') BIGINT(20)"`
	Updatetime        time.Time `json:"updatetime" xorm:"comment('修改时间') TIMESTAMP"`
}
