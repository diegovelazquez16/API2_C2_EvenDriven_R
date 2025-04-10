package controllers


import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"api2/pagos/aplication/usecase"
)

type PagoGetController struct {
	GetPagoUC *usecase.GetPagoUseCase
}

func (c *PagoGetController) GetPagoByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El ID debe estar registrado o ser un numero valido"})
		return
	}

	pago, err := c.GetPagoUC.Execute(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado."})
		return
	}

	ctx.JSON(http.StatusOK, pago)
}
