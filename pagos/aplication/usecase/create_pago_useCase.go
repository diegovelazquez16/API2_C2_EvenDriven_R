package usecase

import (
	"api2/pagos/domain/models"
	"api2/pagos/domain/repository"
)

type CreatePagoUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *CreatePagoUseCase) Execute(pago *models.Pago) error {
	return uc.PagoRepo.Create(pago)
}
