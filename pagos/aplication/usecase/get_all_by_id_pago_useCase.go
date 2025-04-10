package usecase

import (
	"api2/pagos/domain/models"
	"api2/pagos/domain/repository"
)
type GetPagoUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *GetPagoUseCase) Execute(id uint) (*models.Pago, error) {
	return uc.PagoRepo.GetByID(id)
}