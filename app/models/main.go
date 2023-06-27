package models

import (
	"fmt"
	"gin_api/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Video{})
}

func SetInitialConfigs(db *gorm.DB) {
	db.FirstOrCreate(&User{Name: "admin", Password: "123456"})
	db.Create(&User{Name: "teste", Password: "1234567"})

}

func DatabaseConnection(config *config.Config) *gorm.DB {
	dbname := fmt.Sprintf("%v.db", config.DbName)
	DB, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}
	MigrateModels(DB)
	SetInitialConfigs(DB)
	return DB
}
