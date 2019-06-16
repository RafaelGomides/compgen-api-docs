package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20190616_012921 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20190616_012921{}
	m.Created = "20190616_012921"

	migration.Register("User_20190616_012921", m)
}

// Run the migrations
func (m *User_20190616_012921) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE User(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`cpf` varchar(128) NOT NULL,`email` varchar(128) NOT NULL,`password` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *User_20190616_012921) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `User`")
}
