package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Tag 标签模型
type Tag struct {
	ID        int    `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	DeletedAt *time.Time
}

// TableName 表名
func (tag *Tag) TableName() string {
	return "tag"
}

// TagFirst 第一条数据
func TagFirst(db *gorm.DB) {
	tag := new(Tag)

	// 查询条件放在结构体
	// tag.ID = 2

	db.First(tag)
	fmt.Printf("%T: %v", tag, tag)
}
