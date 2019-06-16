package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Permissions_20190616_013156 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Permissions_20190616_013156{}
	m.Created = "20190616_013156"

	migration.Register("Permissions_20190616_013156", m)
}

// Run the migrations
func (m *Permissions_20190616_013156) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Permissions(`id` int(11) NOT NULL AUTO_INCREMENT,`type` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Permissions_20190616_013156) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Permissions`")
}
