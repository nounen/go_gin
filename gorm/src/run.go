package main

import (
	"fmt"
	"go_packages/gorm/src/config"
	"go_packages/gorm/src/model"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// mysql 连接
	db, err := gorm.Open("mysql", config.GetMysqlConf())
	defer db.Close()

	if err != nil {
		fmt.Printf("数据库连接失败: %v", err)
		os.Exit(0)
	}

	// 表名单数全局设置
	// db.SingularTable(true)

	// 查询案例
	model.TagFirst(db)
}
