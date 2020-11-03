package config

import (
	"github.com/taufikhidayatugmbe03/Digitalent-Kominfo_Implementasi-MVC/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB => Initial DB
var DB *gorm.DB

// DBInit => Mysql Connection
func DBInit() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@/simple_bank?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect DB" + err.Error())
	}

	// Automigrate
	db.AutoMigrate(new(model.Account), new(model.Transaction))

	DB = db

	return db
}
