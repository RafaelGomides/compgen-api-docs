package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Document_20190616_013607 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Document_20190616_013607{}
	m.Created = "20190616_013607"

	migration.Register("Document_20190616_013607", m)
}

// Run the migrations
func (m *Document_20190616_013607) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE Document(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`status` varchar(128) NOT NULL,`public_copy` varchar(128) NOT NULL,`private_copy` varchar(128) NOT NULL,`type` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *Document_20190616_013607) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `Document`")
}
