package main

import (
	"fmt"
	"go_packages/gorm/src/config"
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
}
