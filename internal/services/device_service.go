package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type DeviceService interface {
	GetAllDevices() ([]dto.SummaryDeviceDTO, error)
	GetDeviceByID(id uint) (dto.DetailDeviceDTO, error)
	CreateDevice(req dto.CreateDeviceReqDTO) (dto.CreateDeviceResDTO, error)
	UpdateDevice(req dto.UpdateDeviceReqDTO) (dto.UpdateDeviceResDTO, error)
	DeleteDevice(id uint) error
}

type deviceService struct {
	repo repositories.DeviceRepository
}

func NewDeviceService(repo repositories.DeviceRepository) DeviceService {
	return &deviceService{repo}
}

// GET all → Summary DTO
func (s *deviceService) GetAllDevices() ([]dto.SummaryDeviceDTO, error) {
	devices, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SummaryDeviceDTO
	for _, d := range devices {
		res = append(res, dto.SummaryDeviceDTO{
			DeviceID:  d.DeviceID,
			Name:      d.Name,
			IpAddress: d.IpAddress,
			Status:    d.Status,
		})
	}
	return res, nil
}

// GET by ID → Detail DTO
func (s *deviceService) GetDeviceByID(id uint) (dto.DetailDeviceDTO, error) {
	device, err := s.repo.FindByID(id)
	if err != nil {
		return dto.DetailDeviceDTO{}, err
	}

	var airport *dto.SummaryAirportDTO
	if device.Airport != nil {
		airport = &dto.SummaryAirportDTO{
			AirportID: device.Airport.AirportID,
			Name:      device.Airport.Name,
			Code:      device.Airport.Code,
			Address:   device.Airport.Address,
		}
	}

	return dto.DetailDeviceDTO{
		DeviceID:  device.DeviceID,
		Name:      device.Name,
		IpAddress: device.IpAddress,
		Status:    device.Status,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
		Airport:   airport,
	}, nil
}

// POST → Create DTO
func (s *deviceService) CreateDevice(req dto.CreateDeviceReqDTO) (dto.CreateDeviceResDTO, error) {
	device := models.Device{
		AirportID: req.AirportID,
		Name:      req.Name,
		IpAddress: req.IpAddress,
		Status:    req.Status,
	}

	if err := s.repo.Create(&device); err != nil {
		return dto.CreateDeviceResDTO{}, err
	}

	return dto.CreateDeviceResDTO{
		DeviceID:  device.DeviceID,
		Name:      device.Name,
		IpAddress: device.IpAddress,
		Status:    device.Status,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}, nil
}

// PUT/PATCH → Update DTO
func (s *deviceService) UpdateDevice(req dto.UpdateDeviceReqDTO) (dto.UpdateDeviceResDTO, error) {
	device, err := s.repo.FindByID(req.DeviceID)
	if err != nil {
		return dto.UpdateDeviceResDTO{}, err
	}

	if req.AirportID != nil {
		device.AirportID = *req.AirportID
	}
	if req.Name != nil {
		device.Name = *req.Name
	}
	if req.IpAddress != nil {
		device.IpAddress = *req.IpAddress
	}
	if req.Status != nil {
		device.Status = *req.Status
	}

	if err := s.repo.Update(device); err != nil {
		return dto.UpdateDeviceResDTO{}, err
	}

	return dto.UpdateDeviceResDTO{
		DeviceID:  device.DeviceID,
		Name:      device.Name,
		IpAddress: device.IpAddress,
		Status:    device.Status,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}, nil
}

// DELETE
func (s *deviceService) DeleteDevice(id uint) error {
	return s.repo.Delete(id)
}
