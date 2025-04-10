package repository

import (
	"gorm.io/gorm"

	"api2/pagos/domain/models"
)


type PagoRepositoryImpl struct {
	DB *gorm.DB  // Implementación de la interfaz

}
func (r *PagoRepositoryImpl) Create(pago *models.Pago) error {
	return r.DB.Create(pago).Error
}

func (r *PagoRepositoryImpl) GetAll() ([]models.Pago, error) {
	var pagos []models.Pago
	err := r.DB.Find(&pagos).Error
	return pagos, err
}

func (r *PagoRepositoryImpl) GetByID(id uint) (*models.Pago, error) {
	var pago models.Pago
	err := r.DB.First(&pago, id).Error
	return &pago, err
}

func (r *PagoRepositoryImpl) Update(pago *models.Pago) error{
	return r.DB.Save(pago).Error
}

func (r *PagoRepositoryImpl) Delete(id uint) error {
	return r.DB.Delete(&models.Pago{}, id).Error
} 