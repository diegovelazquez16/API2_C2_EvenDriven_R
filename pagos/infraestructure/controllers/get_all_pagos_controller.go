package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"api2/pagos/aplication/usecase"
)
type PagoGetAllController struct {
	GetAllPagosUC *usecase.GetAllPagosUseCase
}
//cambia GetAll a execut o run y lo mismo para todos
func (c *PagoGetAllController) GetAll(ctx *gin.Context) {
	pagos, err := c.GetAllPagosUC.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pagos)
}