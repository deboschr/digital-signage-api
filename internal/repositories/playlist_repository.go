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

	// PlaylistContent
	AddContents(playlistID uint, contents []models.PlaylistContent) error
	UpdateOrders(playlistID uint, contents []models.PlaylistContent) error
	RemoveContents(playlistID uint, contentIDs []uint) error
}

type playlistRepository struct {
	db *gorm.DB
}

func NewPlaylistRepository(db *gorm.DB) PlaylistRepository {
	return &playlistRepository{db}
}

func (r *playlistRepository) FindAll() ([]models.Playlist, error) {
	
	var playlists []models.Playlist
	
	err := r.db.Find(&playlists).Error
	
	return playlists, err
}

func (r *playlistRepository) FindByID(id uint) (*models.Playlist, error) {
	
	var playlist models.Playlist
	
	err := r.db.
		Preload("Airport").
		Preload("Schedules").
		Preload("PlaylistContent.Content").
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

// -----------------------------
// Playlist Content management
// -----------------------------

// Tambah isi playlist
func (r *playlistRepository) AddContents(playlistID uint, contents []models.PlaylistContent) error {
	for _, c := range contents {
		pc := models.PlaylistContent{
			PlaylistID: playlistID,
			ContentID:  c.ContentID,
			Order:      c.Order,
		}
		if err := r.db.Create(&pc).Error; err != nil {
			return err
		}
	}
	return nil
}



// Update urutan konten
func (r *playlistRepository) UpdateOrders(playlistID uint, contents []models.PlaylistContent) error {
	for _, c := range contents {
		if err := r.db.Model(&models.PlaylistContent{}).
			Where("playlist_id = ? AND content_id = ?", playlistID, c.ContentID).
			Update("order", c.Order).Error; err != nil {
			return err
		}
	}
	return nil
}

// Hapus konten dari playlist
func (r *playlistRepository) RemoveContents(playlistID uint, contentIDs []uint) error {
	return r.db.Where("playlist_id = ? AND content_id IN ?", playlistID, contentIDs).
		Delete(&models.PlaylistContent{}).Error
}
