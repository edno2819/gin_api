package models

import (
	"fmt"
	"gin_api/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Video{})
}

func setInitialConfigs(db *gorm.DB) {
	db.FirstOrCreate(&User{Name: "admin", Password: "123456"})
	db.Create(&User{Name: "teste", Password: "1234567"})

}

func DatabaseConnection(config *config.Config) *gorm.DB {
	dbname := fmt.Sprintf("%v.db", config.DbName)
	db, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}
	migrateModels(db)
	setInitialConfigs(db)
	return db
}
