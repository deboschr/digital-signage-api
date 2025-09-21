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
	CreateAirport(req dto.CreateAirportReqDTO) (dto.GetSummaryAirportResDTO, error)
	UpdateAirport(req dto.UpdateAirportReqDTO) (dto.GetSummaryAirportResDTO, error)
	DeleteAirport(id uint) error
}

type airportService struct {
	repository repositories.AirportRepository
}

func NewAirportService(repository repositories.AirportRepository) AirportService {
	return &airportService{repository}
}

func (s *airportService) GetAirports() ([]dto.GetSummaryAirportResDTO, error) {

	airports, err := s.repository.FindAll()

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
			Timezone:  a.Timezone,
		})
	}

	return res, nil
}

func (s *airportService) GetAirport(id uint) (dto.GetDetailAirportResDTO, error) {

	airport, err := s.repository.FindByID(id)

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
			DeviceID:		d.DeviceID,
			Name:				d.Name,
			ApiKey:			d.ApiKey,
			IsConnected:	d.IsConnected,
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
			Title: 		c.Title,
			Type:			c.Type,
			Duration:	c.Duration,
			URL:			utils.BuildContentURL(c.Title),
		}
		contents = append(contents, &content)
	}

	schedules := []*dto.GetSummaryScheduleResDTO{}
	for _, c := range airport.Schedules {
		schedule := dto.GetSummaryScheduleResDTO{
			ScheduleID:		c.ScheduleID,
			StartDate:		c.StartDate,
			EndDate:			c.EndDate,
			StartTime:		c.StartTime,
			EndTime:			c.EndTime,
			RepeatPattern:	c.RepeatPattern,
			IsUrgent:		c.IsUrgent,
		}
		schedules = append(schedules, &schedule)
	}

	return dto.GetDetailAirportResDTO{
		AirportID: airport.AirportID,
		Name:      airport.Name,
		Code:      airport.Code,
		Address:   airport.Address,
		Timezone:  airport.Timezone,
		Users:     users,
		Devices:   devices,
		Contents:  contents,
		Playlists: playlists,
		Schedules: schedules,
	}, nil
}

func (s *airportService) CreateAirport(req dto.CreateAirportReqDTO) (dto.GetSummaryAirportResDTO, error) {
	
	airport := models.Airport{
		Name:			req.Name,
		Code:			req.Code,
		Address:		req.Address,
		Timezone:	req.Timezone,
	}

	if err := s.repository.Create(&airport); err != nil {
		return dto.GetSummaryAirportResDTO{}, err
	}

	return dto.GetSummaryAirportResDTO{
		AirportID:	airport.AirportID,
		Name:			airport.Name,
		Code:			airport.Code,
		Address:		airport.Address,
		Timezone:	airport.Timezone,
	}, nil
}

func (s *airportService) UpdateAirport(req dto.UpdateAirportReqDTO) (dto.GetSummaryAirportResDTO, error) {
	
	airport, err := s.repository.FindByID(req.AirportID)
	if err != nil {
		return dto.GetSummaryAirportResDTO{}, err
	}

	if req.Name != nil {
		airport.Name = *req.Name
	}
	if req.Code != nil {
		airport.Code = *req.Code
	}
	if req.Address != nil {
		airport.Address = *req.Address
	}
	if req.Timezone != nil {
		airport.Timezone = *req.Timezone
	}

	if err := s.repository.Update(airport); err != nil {
		return dto.GetSummaryAirportResDTO{}, err
	}

	return dto.GetSummaryAirportResDTO{
		AirportID:	airport.AirportID,
		Name:			airport.Name,
		Code:			airport.Code,
		Address:		airport.Address,
		Timezone:	airport.Timezone,
	}, nil
}

func (s *airportService) DeleteAirport(id uint) error {
	return s.repository.Delete(id)
}
