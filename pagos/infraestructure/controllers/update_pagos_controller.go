package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api2/pagos/aplication/usecase"
	"api2/pagos/domain/models"
)

type PagoUpdateController struct {
	UpdatePagoUC *usecase.UpdatePagoUseCase
}

func (c *PagoUpdateController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid pago ID",
		})
		return
	}

	var pago models.Pago
	if err := ctx.ShouldBindJSON(&pago); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON payload: " + err.Error(),
		})
		return
	}

	pago.ID = uint(id)

	if err := c.UpdatePagoUC.Execute(&pago); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Pago actualizado correctamente",
		"pago":    pago,
	})
}