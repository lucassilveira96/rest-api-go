package database

import (
	"log"
	"rest-api-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnection() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.User{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.ProductCategory{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.ProductCategory{})
	DB.AutoMigrate(&models.Product{})
	DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Product{})
}
