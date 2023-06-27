package teste

import (
	"fmt"
	"gin_api/app/middlewares"
	"gin_api/app/models"
	"gin_api/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB_test *gorm.DB
)

func databaseConnectionTest(config *config.Config) *gorm.DB {
	dbname := fmt.Sprintf("%v_teste.db", config.DbName)
	DB_test, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}
	models.MigrateModels(DB_test)
	models.SetInitialConfigs(DB_test)
	return DB_test
}

func SetupServerTest() *gin.Engine {
	fmt.Println("ENTROU PARA CRIAR UM NOVO SERVER")
	var conf *config.Config = config.GetConfig()
	config.SetupOutputGinTest(conf)

	db := databaseConnectionTest(conf)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger())

	// Middlewares
	server.Use(middlewares.DatabaseContext(db))

	return server
}

var Server *gin.Engine = SetupServerTest()
