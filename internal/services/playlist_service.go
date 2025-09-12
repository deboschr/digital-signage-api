package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type PlaylistService interface {
	GetAllPlaylists() ([]dto.SummaryPlaylistDTO, error)
	GetPlaylistByID(id uint) (dto.DetailPlaylistDTO, error)
	CreatePlaylist(req dto.CreatePlaylistReqDTO) (dto.CreatePlaylistResDTO, error)
	UpdatePlaylist(req dto.UpdatePlaylistReqDTO) (dto.UpdatePlaylistResDTO, error)
	DeletePlaylist(id uint) error

	// PlaylistContent
	AddContents(req dto.CreatePlaylistContentReqDTO) (dto.DetailPlaylistDTO, error)
	UpdateOrders(req dto.UpdatePlaylistContentReqDTO) (dto.DetailPlaylistDTO, error)
	RemoveContents(req dto.DeletePlaylistContentReqDTO) error
}

type playlistService struct {
	repo repositories.PlaylistRepository
}

func NewPlaylistService(repo repositories.PlaylistRepository) PlaylistService {
	return &playlistService{repo}
}

// GET all → Summary DTO
func (s *playlistService) GetAllPlaylists() ([]dto.SummaryPlaylistDTO, error) {
	playlists, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SummaryPlaylistDTO
	for _, p := range playlists {
		res = append(res, dto.SummaryPlaylistDTO{
			PlaylistID:  p.PlaylistID,
			Name:        p.Name,
			Description: p.Description,
		})
	}
	return res, nil
}

// GET by ID → Detail DTO
func (s *playlistService) GetPlaylistByID(id uint) (dto.DetailPlaylistDTO, error) {
	playlist, err := s.repo.FindByID(id)
	if err != nil {
		return dto.DetailPlaylistDTO{}, err
	}

	var airport *dto.SummaryAirportDTO
	if playlist.Airport != nil {
		airport = &dto.SummaryAirportDTO{
			AirportID: playlist.Airport.AirportID,
			Name:      playlist.Airport.Name,
			Code:      playlist.Airport.Code,
			Address:   playlist.Airport.Address,
		}
	}

	contents := []dto.PlaylistContentItemDTO{}
	for _, pc := range playlist.PlaylistContent {
		if pc.Content != nil {
			contents = append(contents, dto.PlaylistContentItemDTO{
				ContentID: pc.Content.ContentID,
				Title:     pc.Content.Title,
				Order:     pc.Order,
				Duration:  pc.Content.Duration,
			})
		}
	}

	schedules := []dto.SummaryScheduleDTO{}
	for _, sch := range playlist.Schedules {
		schedules = append(schedules, dto.SummaryScheduleDTO{
			ScheduleID:    sch.ScheduleID,
			StartTime:     sch.StartTime,
			EndTime:       sch.EndTime,
			RepeatPattern: sch.RepeatPattern,
		})
	}

	return dto.DetailPlaylistDTO{
		PlaylistID:  playlist.PlaylistID,
		Name:        playlist.Name,
		Description: playlist.Description,
		CreatedAt:   playlist.CreatedAt,
		UpdatedAt:   playlist.UpdatedAt,
		Airport:     airport,
		Contents:    contents,
		Schedules:   schedules,
	}, nil
}

// POST → Create DTO
func (s *playlistService) CreatePlaylist(req dto.CreatePlaylistReqDTO) (dto.CreatePlaylistResDTO, error) {
	playlist := models.Playlist{
		AirportID:   req.AirportID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Create(&playlist); err != nil {
		return dto.CreatePlaylistResDTO{}, err
	}

	return dto.CreatePlaylistResDTO{
		PlaylistID:  playlist.PlaylistID,
		Name:        playlist.Name,
		Description: playlist.Description,
		CreatedAt:   playlist.CreatedAt,
		UpdatedAt:   playlist.UpdatedAt,
	}, nil
}

// PUT/PATCH → Update DTO
func (s *playlistService) UpdatePlaylist(req dto.UpdatePlaylistReqDTO) (dto.UpdatePlaylistResDTO, error) {
	playlist, err := s.repo.FindByID(req.PlaylistID)
	if err != nil {
		return dto.UpdatePlaylistResDTO{}, err
	}

	if req.AirportID != nil {
		playlist.AirportID = *req.AirportID
	}
	if req.Name != nil {
		playlist.Name = *req.Name
	}
	if req.Description != nil {
		playlist.Description = *req.Description
	}

	if err := s.repo.Update(playlist); err != nil {
		return dto.UpdatePlaylistResDTO{}, err
	}

	return dto.UpdatePlaylistResDTO{
		PlaylistID:  playlist.PlaylistID,
		Name:        playlist.Name,
		Description: playlist.Description,
		CreatedAt:   playlist.CreatedAt,
		UpdatedAt:   playlist.UpdatedAt,
	}, nil
}

// DELETE
func (s *playlistService) DeletePlaylist(id uint) error {
	return s.repo.Delete(id)
}

// -----------------------------
// PlaylistContent management
// -----------------------------

// POST /playlists/content
func (s *playlistService) AddContents(req dto.CreatePlaylistContentReqDTO) (dto.DetailPlaylistDTO, error) {
	var contents []models.PlaylistContent
	for _, c := range req.Contents {
		contents = append(contents, models.PlaylistContent{
			PlaylistID: req.PlaylistID,
			ContentID:  c.ContentID,
			Order:      c.Order,
		})
	}
	if err := s.repo.AddContents(req.PlaylistID, contents); err != nil {
		return dto.DetailPlaylistDTO{}, err
	}
	// reload playlist
	return s.GetPlaylistByID(req.PlaylistID)
}



// PATCH /playlists/content
func (s *playlistService) UpdateOrders(req dto.UpdatePlaylistContentReqDTO) (dto.DetailPlaylistDTO, error) {
	var contents []models.PlaylistContent
	for _, c := range req.Contents {
		contents = append(contents, models.PlaylistContent{
			PlaylistID: req.PlaylistID,
			ContentID:  c.ContentID,
			Order:      c.Order,
		})
	}
	if err := s.repo.UpdateOrders(req.PlaylistID, contents); err != nil {
		return dto.DetailPlaylistDTO{}, err
	}
	// reload playlist
	return s.GetPlaylistByID(req.PlaylistID)
}

// DELETE /playlists/content
func (s *playlistService) RemoveContents(req dto.DeletePlaylistContentReqDTO) error {
	return s.repo.RemoveContents(req.PlaylistID, req.ContentIDs)
}
