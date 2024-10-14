package config

import (
	"animal-rest/models"
	"animal-rest/util"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type DBConfig struct {
	DB_Host string
	DB_Name string
	DB_User string
	DB_Pass string
	DB_Port string
}

func InitDB() {
	dbConfig := DBConfig{
		DB_Host: util.GetEnv("DB_HOST", "localhost"),
		DB_Name: util.GetEnv("DB_NAME", "animal"),
		DB_User: util.GetEnv("DB_USER", "root"),
		DB_Pass: util.GetEnv("DB_PASS", ""),
		DB_Port: util.GetEnv("DB_PORT", "3306"),
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.DB_User,
		dbConfig.DB_Pass,
		dbConfig.DB_Host,
		dbConfig.DB_Port,
		dbConfig.DB_Name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitMigration() {
	DB.AutoMigrate(&models.Animal{})
}
