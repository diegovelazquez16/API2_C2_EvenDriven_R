package main
import (
	"api2/core"
	"api2/launch"
	"api2/pagos/aplication/usecase"
	"api2/pagos/domain/repository" // Asegúrate de importar el repositorio
	"api2/pagos/infraestructure/messaging"  // Asegúrate de importar RabbitMQ
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitializeApp()

	pagoRepo := &repository.PagoRepositoryImpl{DB: core.GetDB()} 

	pagoUseCase := &usecase.CreatePagoUseCase{PagoRepo: pagoRepo} 

	pedidoConsumer, err := messaging.NewPedidoConsumer(pagoUseCase)
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer pedidoConsumer.Close()

	go pedidoConsumer.StartConsuming()

	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	launch.RegisterRoutes(app)

	log.Println("API corriendo en http://localhost:8081")
	if err := app.Run(":8081"); err != nil {
		log.Fatalf("Error al correr el servidor: %v", err)
	}
}
