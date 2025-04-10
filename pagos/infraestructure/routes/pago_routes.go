package routes

import (
	
	"api2/pagos/infraestructure/controllers"
	"github.com/gin-gonic/gin"

)

func PagoRoutes(
	router *gin.Engine,
	createController *controllers.PagoCreateController,
	getAllController *controllers.PagoGetAllController,
	updateController *controllers.PagoUpdateController,
	deleteController *controllers.PagoDeleteController,
	getController *controllers.PagoGetController,

) {
	group := router.Group("/pagos")
	{
		group.POST("", createController.Create)          
		group.GET("", getAllController.GetAll)  
		group.GET("/:id", getController.GetPagoByID)   
		group.PUT("/:id", updateController.Update)    
		group.DELETE("/:id", deleteController.Delete )   
    
	}
}