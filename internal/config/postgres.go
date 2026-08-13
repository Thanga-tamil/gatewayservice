package config

import (
	log "gateway/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Psql *gorm.DB

type settings struct {
	Licensekey string
	IV string
}

var ChatSettings *settings

func NewPostgreSql(){
	dsn := "host=localhost user=postgres password=root dbname=chat port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Err while opening connection to postgresql: ", err.Error())
	}

	if err := db.Exec("SET SEARCH_PATH TO chat;").Error; err != nil {
		log.Fatalf(err.Error())
	}

	loadChatSettings(db)

	Psql = db // bollocks
	log.Info("✅ PostgreSQL connected successfully")
}

func loadChatSettings(db *gorm.DB){
	if err := db.Raw(`SELECT licensekey, iv FROM chat_settings;`).Scan(&ChatSettings).Error; err != nil {
		log.Fatalf(err.Error())
	}
}
