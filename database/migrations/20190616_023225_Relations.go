package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Relations_20190616_023225 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Relations_20190616_023225{}
	m.Created = "20190616_023225"

	migration.Register("Relations_20190616_023225", m)
}

// Run the migrations
func (m *Relations_20190616_023225) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE `Document` ADD COLUMN `user` INT(11) NOT NULL AFTER `type`, ADD COLUMN `service_order` INT(11) NOT NULL AFTER `user`,	DROP PRIMARY KEY,	ADD PRIMARY KEY (`id`, `user`, `service_order`), ADD INDEX `fk_Document_User1_idx` (`user` ASC) VISIBLE,	ADD INDEX `fk_Document_ServiceOrder1_idx` (`service_order` ASC) VISIBLE;")
	m.SQL("ALTER TABLE `ServiceOrder` ADD COLUMN `service` INT(11) NOT NULL AFTER `status`,	DROP PRIMARY KEY,	ADD PRIMARY KEY (`id`, `service`), ADD INDEX `fk_ServiceOrder_Services1_idx` (`service` ASC) VISIBLE;")
	m.SQL("ALTER TABLE `User` ADD COLUMN `permission` INT(11) NOT NULL AFTER `password`, DROP PRIMARY KEY,	ADD PRIMARY KEY (`id`, `permission`),	ADD INDEX `fk_User_Permissions_idx` (`permission` ASC) VISIBLE;")
	m.SQL("CREATE TABLE IF NOT EXISTS `ServiceOrderUsers` (`service_order` INT(11) NOT NULL,	`user` INT(11) NOT NULL, `created_at` DATETIME NULL DEFAULT NOW(),	PRIMARY KEY (`service_order`, `user`), INDEX `fk_ServiceOrder_has_User_User1_idx` (`user` ASC) VISIBLE, INDEX `fk_ServiceOrder_has_User_ServiceOrder1_idx` (`service_order` ASC) VISIBLE,	CONSTRAINT `fk_ServiceOrder_has_User_ServiceOrder1`	FOREIGN KEY (`service_order`)	REFERENCES `ServiceOrder` (`id`)	ON DELETE NO ACTION	ON UPDATE NO ACTION, CONSTRAINT `fk_ServiceOrder_has_User_User1`	FOREIGN KEY (`user`) REFERENCES `User` (`id`)	ON DELETE NO ACTION	ON UPDATE NO ACTION) ENGINE = InnoDB	DEFAULT CHARACTER SET = utf8mb4	COLLATE = utf8mb4_0900_ai_ci;")
	m.SQL("ALTER TABLE `Document` ADD CONSTRAINT `fk_Document_User1` FOREIGN KEY (`user`) REFERENCES `User` (`id`)	ON DELETE NO ACTION	ON UPDATE NO ACTION, ADD CONSTRAINT `fk_Document_ServiceOrder1` FOREIGN KEY (`service_order`)	REFERENCES `ServiceOrder` (`id`)	ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE `ServiceOrder` ADD CONSTRAINT `fk_ServiceOrder_Services1` FOREIGN KEY (`service`)	REFERENCES `Services` (`id`)	ON DELETE NO ACTION ON UPDATE NO ACTION;")
	m.SQL("ALTER TABLE `User` ADD CONSTRAINT `fk_User_Permissions` FOREIGN KEY (`permission`) REFERENCES `Permissions` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION;")
}

// Reverse the migrations
func (m *Relations_20190616_023225) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
