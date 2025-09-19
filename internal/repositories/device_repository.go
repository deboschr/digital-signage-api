package repositories

import (
	"digital_signage_api/internal/models"

	"gorm.io/gorm"
)

type DeviceRepository interface {
	FindAll() ([]models.Device, error)
	FindByID(id uint) (*models.Device, error)
	FindByAirport(airportID uint)([]models.Device, error)
	Create(device *models.Device) error
	Update(device *models.Device) error
	Delete(id uint) error
}

type deviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) DeviceRepository {
	return &deviceRepository{db}
}

func (r *deviceRepository) FindAll() ([]models.Device, error) {
	var devices []models.Device
	err := r.db.Find(&devices).Error
	return devices, err
}

func (r *deviceRepository) FindByAirport(airportID uint) ([]models.Device, error) {
	var devices []models.Device
	err := r.db.Preload("Airport").
		Where("airport_id = ?", airportID).
		Find(&devices).Error
	return devices, err
}


func (r *deviceRepository) FindByID(id uint) (*models.Device, error) {
	var device models.Device
	err := r.db.Preload("Airport").
		First(&device, "device_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (r *deviceRepository) Create(device *models.Device) error {
	return r.db.Create(device).Error
}

func (r *deviceRepository) Update(device *models.Device) error {
	return r.db.Save(device).Error
}

func (r *deviceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Device{}, "device_id = ?", id).Error
}
