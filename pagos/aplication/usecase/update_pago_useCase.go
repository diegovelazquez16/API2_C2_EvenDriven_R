package usecase

import (
	"api2/pagos/domain/models"
	"api2/pagos/domain/repository"
)

type UpdatePagoUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *UpdatePagoUseCase) Execute(pago *models.Pago) error {
	return uc.PagoRepo.Update(pago)
}