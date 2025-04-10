package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"api2/pagos/aplication/usecase"
)

type PagoDeleteController struct {
	DeletePagoUC *usecase.DeletePagoUseCase
}

func (c *PagoDeleteController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Pago ID"})
		return
	}

	err = c.DeletePagoUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pago deleted successfully"})
}