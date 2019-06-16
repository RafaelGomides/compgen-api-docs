package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type ServiceOrder_20190616_013445 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ServiceOrder_20190616_013445{}
	m.Created = "20190616_013445"

	migration.Register("ServiceOrder_20190616_013445", m)
}

// Run the migrations
func (m *ServiceOrder_20190616_013445) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE ServiceOrder(`id` int(11) NOT NULL AUTO_INCREMENT,`status` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *ServiceOrder_20190616_013445) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `ServiceOrder`")
}
