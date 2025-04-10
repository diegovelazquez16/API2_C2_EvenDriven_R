package repository

import (
	"api2/pagos/domain/models"
)

type IPagoRepository interface {
	Create(pago *models.Pago) error
	GetAll() ([]models.Pago, error)
	GetByID(id uint) (*models.Pago, error)
	Update(pago *models.Pago) error// Interfaz para conectar el modelo con el repositorio
	Delete(id uint) error
}
