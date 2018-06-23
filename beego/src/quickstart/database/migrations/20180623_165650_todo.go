package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Todo_20180623_165650 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Todo_20180623_165650{}
	m.Created = "20180623_165650"

	migration.Register("Todo_20180623_165650", m)
}

// Run the migrations
func (m *Todo_20180623_165650) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE todo(`id` int(11) NOT NULL AUTO_INCREMENT,`title` varchar(128) NOT NULL,`sort` int(11) DEFAULT NULL,`status` int(11) DEFAULT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Todo_20180623_165650) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `todo`")
}
