package services

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type DeviceService interface {
	GetAllDevices() ([]models.Device, error)
	GetDeviceByID(id uint) (*models.Device, error)
	CreateDevice(device *models.Device) error
	UpdateDevice(device *models.Device) error
	DeleteDevice(id uint) error
}

type deviceService struct {
	repo repositories.DeviceRepository
}

func NewDeviceService(repo repositories.DeviceRepository) DeviceService {
	return &deviceService{repo}
}

func (s *deviceService) GetAllDevices() ([]models.Device, error) {
	return s.repo.FindAll()
}

func (s *deviceService) GetDeviceByID(id uint) (*models.Device, error) {
	return s.repo.FindByID(id)
}

func (s *deviceService) CreateDevice(device *models.Device) error {
	return s.repo.Create(device)
}

func (s *deviceService) UpdateDevice(device *models.Device) error {
	return s.repo.Update(device)
}

func (s *deviceService) DeleteDevice(id uint) error {
	return s.repo.Delete(id)
}
