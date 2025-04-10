package usecase

import (
	"api2/pagos/domain/models"
	"api2/pagos/domain/repository"
)
type GetAllPagosUseCase struct {
	PagoRepo repository.IPagoRepository
}

func (uc *GetAllPagosUseCase) Execute() ([]models.Pago, error) {
	return uc.PagoRepo.GetAll()
}

