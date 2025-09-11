package repositories

import (
	"digital_signage_api/internal/models"

	"gorm.io/gorm"
)

type PlaylistRepository interface {
	FindAll() ([]models.Playlist, error)
	FindByID(id uint) (*models.Playlist, error)
	Create(playlist *models.Playlist) error
	Update(playlist *models.Playlist) error
	Delete(id uint) error
}

type playlistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) PlaylistRepository {
	return &playlistRepository{db}
}

func (r *playlistRepository) FindAll() ([]models.Playlist, error) {
	var playlists []models.Playlist
	err := r.db.
		Preload("Airport").
		Preload("Schedules").
		Preload("Contents").
		Find(&playlists).Error
	return playlists, err
}

func (r *playlistRepository) FindByID(id uint) (*models.Playlist, error) {
	var playlist models.Playlist
	err := r.db.
		Preload("Airport").
		Preload("Schedules").
		Preload("Contents").
		First(&playlist, "playlist_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &playlist, nil
}

func (r *playlistRepository) Create(playlist *models.Playlist) error {
	return r.db.Create(playlist).Error
}

func (r *playlistRepository) Update(playlist *models.Playlist) error {
	return r.db.Save(playlist).Error
}

func (r *playlistRepository) Delete(id uint) error {
	return r.db.Delete(&models.Playlist{}, "playlist_id = ?", id).Error
}
