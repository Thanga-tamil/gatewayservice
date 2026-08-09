package config

import (
	log "gateway/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Psql *gorm.DB

func NewPostgreSql(){
	dsn := "host=localhost user=postgres password=root dbname=chat port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Error("Err while opening connection to postgresql: ", err.Error()); panic(err)
	}

	if err := db.Exec("SET SEARCH_PATH TO chat;").Error; err != nil {
		panic(err)
	}

	Psql = db // bollocks
	log.Info("PostgreSQL connected successfully")
}
