package core

import (
	"log"

	"github.com/joho/godotenv"
)

// Aqui se cargan las configuraciones desde el env
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando el archivo .env (usando variables de entorno del sistema)")
	}
}
