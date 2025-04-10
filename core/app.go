package core

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitializeApp() *gin.Engine {
	LoadConfig()

	DatabaseConnection()

	app := gin.Default()

	log.Println("Aplicación inicializada correctamente.")

	return app
}
