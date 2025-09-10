package repositories

import (
	"digital_signage_api/internal/db"
	"digital_signage_api/internal/models"
)

type DeviceRepository struct{}

func NewDeviceRepository() *DeviceRepository {
    return &DeviceRepository{}
}

func (r *DeviceRepository) FindAll() ([]models.Device, error) {
    var devices []models.Device
    result := db.DB.Find(&devices)
    return devices, result.Error
}

func (r *DeviceRepository) Create(device *models.Device) error {
    return db.DB.Create(device).Error
}
