package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type AirportService interface {
	GetAllAirports() ([]dto.SummaryAirportDTO, error)
	GetAirportByID(id uint) (dto.DetailAirportDTO, error)
	CreateAirport(req dto.CreateAirportReqDTO) (dto.CreateAirportResDTO, error)
	UpdateAirport(req dto.UpdateAirportReqDTO) (dto.UpdateAirportResDTO, error)
	DeleteAirport(id uint) error
}

type airportService struct {
	repo repositories.AirportRepository
}

func NewAirportService(repo repositories.AirportRepository) AirportService {
	return &airportService{repo}
}

// GET all → Summary DTO
func (s *airportService) GetAllAirports() ([]dto.SummaryAirportDTO, error) {
	airports, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SummaryAirportDTO
	for _, a := range airports {
		res = append(res, dto.SummaryAirportDTO{
			AirportID: a.AirportID,
			Name:      a.Name,
			Code:      a.Code,
			Address:   a.Address,
		})
	}
	return res, nil
}

// GET by ID → Detail DTO
func (s *airportService) GetAirportByID(id uint) (dto.DetailAirportDTO, error) {
	airport, err := s.repo.FindByID(id)
	if err != nil {
		return dto.DetailAirportDTO{}, err
	}

	// mapping relasi
	users := []dto.SummaryUserDTO{}
	for _, u := range airport.Users {
		users = append(users, dto.SummaryUserDTO{
			UserID:   u.UserID,
			Username: u.Username,
			Role:     u.Role,
		})
	}
	devices := []dto.SummaryDeviceDTO{}
	for _, d := range airport.Devices {
		devices = append(devices, dto.SummaryDeviceDTO{
			DeviceID:  d.DeviceID,
			Name:      d.Name,
			IpAddress: d.IpAddress,
			Status:    d.Status,
		})
	}
	playlists := []dto.SummaryPlaylistDTO{}
	for _, p := range airport.Playlists {
		playlists = append(playlists, dto.SummaryPlaylistDTO{
			PlaylistID:  p.PlaylistID,
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return dto.DetailAirportDTO{
		AirportID: airport.AirportID,
		Name:      airport.Name,
		Code:      airport.Code,
		Address:   airport.Address,
		CreatedAt: airport.CreatedAt,
		UpdatedAt: airport.UpdatedAt,
		Users:     users,
		Devices:   devices,
		Playlists: playlists,
	}, nil
}

// POST → Create DTO
func (s *airportService) CreateAirport(req dto.CreateAirportReqDTO) (dto.CreateAirportResDTO, error) {
	airport := models.Airport{
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
	}

	if err := s.repo.Create(&airport); err != nil {
		return dto.CreateAirportResDTO{}, err
	}

	return dto.CreateAirportResDTO{
		AirportID: airport.AirportID,
		Name:      airport.Name,
		Code:      airport.Code,
		Address:   airport.Address,
		CreatedAt: airport.CreatedAt,
		UpdatedAt: airport.UpdatedAt,
	}, nil
}

// PUT/PATCH → Update DTO
func (s *airportService) UpdateAirport(req dto.UpdateAirportReqDTO) (dto.UpdateAirportResDTO, error) {
	// ambil data lama
	airport, err := s.repo.FindByID(req.AirportID)
	if err != nil {
		return dto.UpdateAirportResDTO{}, err
	}

	// update field kalau ada
	if req.Name != nil {
		airport.Name = *req.Name
	}
	if req.Code != nil {
		airport.Code = *req.Code
	}
	if req.Address != nil {
		airport.Address = *req.Address
	}

	if err := s.repo.Update(airport); err != nil {
		return dto.UpdateAirportResDTO{}, err
	}

	return dto.UpdateAirportResDTO{
		AirportID: airport.AirportID,
		Name:      airport.Name,
		Code:      airport.Code,
		Address:   airport.Address,
		CreatedAt: airport.CreatedAt,
		UpdatedAt: airport.UpdatedAt,
	}, nil
}

// DELETE
func (s *airportService) DeleteAirport(id uint) error {
	return s.repo.Delete(id)
}
