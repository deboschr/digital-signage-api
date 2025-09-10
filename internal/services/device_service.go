package services

import (
    "digital_signage_api/internal/models"
    "digital_signage_api/internal/repositories"
)

type DeviceService struct {
    repo *repositories.DeviceRepository
}

func NewDeviceService(repo *repositories.DeviceRepository) *DeviceService {
    return &DeviceService{repo: repo}
}

func (s *DeviceService) GetDevices() ([]models.Device, error) {
    return s.repo.FindAll()
}

func (s *DeviceService) AddDevice(device *models.Device) error {
    return s.repo.Create(device)
}
