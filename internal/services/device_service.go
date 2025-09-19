package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type DeviceService interface {
	GetDevices() ([]dto.GetSummaryDeviceResDTO, error)
	GetDevice(id uint) (dto.GetDetailDeviceResDTO, error)
	CreateDevice(req dto.CreateDeviceReqDTO) (dto.GetSummaryDeviceResDTO, error)
	UpdateDevice(req dto.UpdateDeviceReqDTO) (dto.GetSummaryDeviceResDTO, error)
	DeleteDevice(id uint) error
}

type deviceService struct {
	repo repositories.DeviceRepository
}

func NewDeviceService(repo repositories.DeviceRepository) DeviceService {
	return &deviceService{repo}
}

// GET all → Summary DTO
func (s *deviceService) GetDevices() ([]dto.GetSummaryDeviceResDTO, error) {
	devices, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.GetSummaryDeviceResDTO
	for _, d := range devices {
		res = append(res, dto.GetSummaryDeviceResDTO{
			DeviceID:  d.DeviceID,
			Name:      d.Name,
			IsConnected:    d.IsConnected,
			Airport: d.Airport.Name,
		})
	}
	return res, nil
}

func (s *deviceService) GetDevicesByAirport(airportID uint) ([]dto.GetSummaryDeviceResDTO, error) {
	devices, err := s.repo.FindByAirport(airportID)
	if err != nil {
		return nil, err
	}

	var res []dto.GetSummaryDeviceResDTO
	for _, d := range devices {
		res = append(res, dto.GetSummaryDeviceResDTO{
			DeviceID:    d.DeviceID,
			Name:        d.Name,
			IsConnected: d.IsConnected,
			Airport:     d.Airport.Name,
		})
	}
	return res, nil
}


// GET by ID → Detail DTO
func (s *deviceService) GetDevice(id uint) (dto.GetDetailDeviceResDTO, error) {
	device, err := s.repo.FindByID(id)
	if err != nil {
		return dto.GetDetailDeviceResDTO{}, err
	}

	var airport *dto.GetSummaryAirportResDTO
	if device.Airport != nil {
		airport = &dto.GetSummaryAirportResDTO{
			AirportID: device.Airport.AirportID,
			Name:      device.Airport.Name,
			Code:      device.Airport.Code,
			Address:   device.Airport.Address,
		}
	}

	return dto.GetDetailDeviceResDTO{
		DeviceID:  device.DeviceID,
		Name:      device.Name,
		Airport:   airport,
	}, nil
}

// POST → Create DTO
func (s *deviceService) CreateDevice(req dto.CreateDeviceReqDTO) (dto.GetSummaryDeviceResDTO, error) {
	device := models.Device{
		AirportID: req.AirportID,
		Name:      req.Name,
		IpAddress: req.IpAddress,
		Status:    req.Status,
	}

	if err := s.repo.Create(&device); err != nil {
		return dto.GetSummaryDeviceResDTO{}, err
	}

	return dto.GetSummaryDeviceResDTO{
		DeviceID:  device.DeviceID,
		Name:      device.Name,
		IpAddress: device.IpAddress,
		Status:    device.Status,
		CreatedAt: device.CreatedAt,
		UpdatedAt: device.UpdatedAt,
	}, nil
}

// PUT/PATCH → Update DTO
func (s *deviceService) UpdateDevice(req dto.UpdateDeviceReqDTO) (dto.GetSummaryDeviceResDTO, error) {
	device, err := s.repo.FindByID(req.DeviceID)
	if err != nil {
		return dto.GetSummaryDeviceResDTO{}, err
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
		return dto.GetSummaryDeviceResDTO{}, err
	}

	return dto.GetSummaryDeviceResDTO{
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
