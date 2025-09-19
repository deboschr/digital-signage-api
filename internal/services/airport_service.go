package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/utils"
)

type AirportService interface {
	GetAirports() ([]dto.GetSummaryAirportResDTO, error)
	GetAirport(id uint) (dto.GetDetailAirportResDTO, error)
	CreateAirport(req dto.CreateAirportReqDTO) (dto.GetDetailAirportResDTO, error)
	UpdateAirport(req dto.UpdateAirportReqDTO) (dto.GetDetailAirportResDTO, error)
	DeleteAirport(id uint) error
}

type airportService struct {
	repo repositories.AirportRepository
}

func NewAirportService(repo repositories.AirportRepository) AirportService {
	return &airportService{repo}
}

func (s *airportService) GetAirports() ([]dto.GetSummaryAirportResDTO, error) {

	airports, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var res []dto.GetSummaryAirportResDTO

	for _, a := range airports {
		res = append(res, dto.GetSummaryAirportResDTO{
			AirportID: a.AirportID,
			Name:      a.Name,
			Code:      a.Code,
			Address:   a.Address,
		})
	}

	return res, nil
}

func (s *airportService) GetAirport(id uint) (dto.GetDetailAirportResDTO, error) {

	airport, err := s.repo.FindByID(id)

	if err != nil {
		return dto.GetDetailAirportResDTO{}, err
	}


	users := []*dto.GetSummaryUserResDTO{}
	for _, u := range airport.Users {
    	user := dto.GetSummaryUserResDTO{
        	UserID:   u.UserID,
        	Username: u.Username,
        	Role:     u.Role,
    	}
    	users = append(users, &user)
	}


	devices := []*dto.GetSummaryDeviceResDTO{}
	for _, d := range airport.Devices {
		device := dto.GetSummaryDeviceResDTO{
			DeviceID:  d.DeviceID,
			Name:      d.Name,
			IsConnected:    d.IsConnected,
		}
		devices = append(devices, &device)
	}

	playlists := []*dto.GetSummaryPlaylistResDTO{}
	for _, p := range airport.Playlists {
		playlist := dto.GetSummaryPlaylistResDTO{
			PlaylistID:  p.PlaylistID,
			Name:        p.Name,
			Description: p.Description,
		}
		playlists = append(playlists, &playlist)
	}

	contents := []*dto.GetSummaryContentResDTO{}
	for _, c := range airport.Contents {
		content := dto.GetSummaryContentResDTO{
			ContentID:  c.ContentID,
			Title: c.Title,
			Type:        c.Type,
			Duration:        c.Duration,
			URL:        utils.BuildContentURL(c.Title),
		}
		contents = append(contents, &content)
	}

	return dto.GetDetailAirportResDTO{
		AirportID: airport.AirportID,
		Name:      airport.Name,
		Code:      airport.Code,
		Address:   airport.Address,
		Users:     users,
		Devices:   devices,
		Playlists: playlists,
		Contents: contents,
	}, nil
}

func (s *airportService) CreateAirport(req dto.CreateAirportReqDTO) (dto.GetDetailAirportResDTO, error) {
	
	airport := models.Airport{
		Name:    req.Name,
		Code:    req.Code,
		Address: req.Address,
	}

	if err := s.repo.Create(&airport); err != nil {
		return dto.GetDetailAirportResDTO{}, err
	}

	return dto.GetDetailAirportResDTO{
		AirportID: airport.AirportID,
		Name:      airport.Name,
		Code:      airport.Code,
		Address:   airport.Address,
	}, nil
}

func (s *airportService) UpdateAirport(req dto.UpdateAirportReqDTO) (dto.GetDetailAirportResDTO, error) {
	// ambil data lama
	airport, err := s.repo.FindByID(req.AirportID)
	if err != nil {
		return dto.GetDetailAirportResDTO{}, err
	}

	// update field kalau ada
	if req.Name != nil {
		airport.Name = *req.Name
	}
	if req.Code != nil {
		airport.Code = *req.Code
	}
	if req.Address != nil {
		airport.Address = req.Address
	}

	if err := s.repo.Update(airport); err != nil {
		return dto.GetDetailAirportResDTO{}, err
	}

	return dto.GetDetailAirportResDTO{
		AirportID: airport.AirportID,
		Name:      airport.Name,
		Code:      airport.Code,
		Address:   airport.Address,
	}, nil
}

func (s *airportService) DeleteAirport(id uint) error {
	return s.repo.Delete(id)
}
