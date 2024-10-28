package db

import (
	"fmt"

	"com.mensssy.LabMS/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func DBConnect() {
	SqlConnect()
	MongoConnect()
}

func MongoConnect() {
}

func SqlConnect() {
	dsn := conf.SqlUserName + ":" + conf.SqlPassword + conf.SqlProtocol +
		"(" + conf.SqlHost + ":" + conf.SqlPort + ")/" +
		conf.SqlName + "?" + conf.SqlParam

	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect mysql")
	}

	SqlDB = db

	//sql数据库自动迁移
	SqlMigrate()
}
