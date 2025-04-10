package launch

import (
	"api2/core"
	pagoUsecase "api2/pagos/aplication/usecase"
	pagoRepo "api2/pagos/domain/repository"
	pagoControllers "api2/pagos/infraestructure/controllers"
	pagoRoutes "api2/pagos/infraestructure/routes"
	"github.com/gin-gonic/gin"
)

func RegisterPagoModule(router *gin.Engine) {
	pagoRepo := &pagoRepo.PagoRepositoryImpl{DB: core.GetDB()}

	createPagoUC := &pagoUsecase.CreatePagoUseCase{PagoRepo: pagoRepo}
	getAllPagosUC := &pagoUsecase.GetAllPagosUseCase{PagoRepo: pagoRepo}
	getPagoUC := &pagoUsecase.GetPagoUseCase{PagoRepo: pagoRepo}
	updatePagoUC := &pagoUsecase.UpdatePagoUseCase{PagoRepo: pagoRepo}
	deletePagoUC := &pagoUsecase.DeletePagoUseCase{PagoRepo: pagoRepo}

	pagoCreateController := &pagoControllers.PagoCreateController{CreatePagoUC: createPagoUC}
	pagoGetAllController := &pagoControllers.PagoGetAllController{GetAllPagosUC: getAllPagosUC}
	pagoGetController := &pagoControllers.PagoGetController{GetPagoUC: getPagoUC}
	pagoUpdateController := &pagoControllers.PagoUpdateController{UpdatePagoUC: updatePagoUC}
	pagoDeleteController := &pagoControllers.PagoDeleteController{DeletePagoUC: deletePagoUC}

	pagoRoutes.PagoRoutes(router, pagoCreateController, pagoGetAllController, pagoUpdateController, pagoDeleteController, pagoGetController)
}