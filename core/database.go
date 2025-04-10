package core

import (
	"fmt"
	"log"
	"os"



	pagoModels "api2/pagos/domain/models"


	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func DatabaseConnection() {
	sqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error conectándose a la base de datos: %v", err)
	}

	DB = db
	log.Println("Conexión a la base de datos exitosa.")
	err = DB.AutoMigrate(&pagoModels.Pago{})
	if err != nil {
		log.Fatalf("Error aplicando migración: %v", err)
	}
	log.Println("Migración aplicada exitosamente.")
}

func GetDB() *gorm.DB {
	return DB
}
