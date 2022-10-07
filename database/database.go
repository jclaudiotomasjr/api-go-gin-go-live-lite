package database

import (
	"log"

	"github.com/jclaudiotomasjr/api-gin-go-live/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBD() {

	DB, err = gorm.Open(sqlite.Open("data/db_api_go"), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com Banco de dados!")
	}

	DB.AutoMigrate(&models.Score{})
}
