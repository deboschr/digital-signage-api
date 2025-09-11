package services

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type PlaylistService interface {
	GetAllPlaylists() ([]models.Playlist, error)
	GetPlaylistByID(id uint) (*models.Playlist, error)
	CreatePlaylist(playlist *models.Playlist) error
	UpdatePlaylist(playlist *models.Playlist) error
	DeletePlaylist(id uint) error
}

type playlistService struct {
	repo repositories.PlaylistRepository
}

func NewPlaylistService(repo repositories.PlaylistRepository) PlaylistService {
	return &playlistService{repo}
}

func (s *playlistService) GetAllPlaylists() ([]models.Playlist, error) {
	return s.repo.FindAll()
}

func (s *playlistService) GetPlaylistByID(id uint) (*models.Playlist, error) {
	return s.repo.FindByID(id)
}

func (s *playlistService) CreatePlaylist(playlist *models.Playlist) error {
	return s.repo.Create(playlist)
}

func (s *playlistService) UpdatePlaylist(playlist *models.Playlist) error {
	return s.repo.Update(playlist)
}

func (s *playlistService) DeletePlaylist(id uint) error {
	return s.repo.Delete(id)
}
