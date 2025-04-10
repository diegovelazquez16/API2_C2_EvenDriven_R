package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"api2/pagos/aplication/usecase"
	"api2/pagos/domain/models"
)

type PagoCreateController struct {
	CreatePagoUC *usecase.CreatePagoUseCase
}

func (c *PagoCreateController) Create(ctx *gin.Context) {
	var pago models.Pago
	if err := ctx.ShouldBindJSON(&pago); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.CreatePagoUC.Execute(&pago)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Pago creado de forma exitosa",
		"pago": pago,
	})
}
