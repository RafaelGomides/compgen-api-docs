package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Services_20190616_013344 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Services_20190616_013344{}
	m.Created = "20190616_013344"

	migration.Register("Services_20190616_013344", m)
}

// Run the migrations
func (m *Services_20190616_013344) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Services(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Services_20190616_013344) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Services`")
}
