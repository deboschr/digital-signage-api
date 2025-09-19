package repositories

import (
	"digital_signage_api/internal/models"

	"gorm.io/gorm"
)

type ContentRepository interface {
	FindAll() ([]models.Content, error)
	FindByID(id uint) (*models.Content, error)
	Create(content *models.Content) error
	Update(content *models.Content) error
	Delete(id uint) error
}

type contentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) ContentRepository {
	return &contentRepository{db}
}

func (r *contentRepository) FindAll() ([]models.Content, error) {
	
	var contents []models.Content
	
	err := r.db.
		Preload("Playlists").
		Preload("Playlists.Airport").
		Find(&contents).Error
	
	return contents, err
}

func (r *contentRepository) FindByID(id uint) (*models.Content, error) {
	
	var content models.Content
	
	err := r.db.
		Preload("Playlists").
		Preload("Playlists.Airport").
		First(&content, "content_id = ?", id).Error
	
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *contentRepository) Create(content *models.Content) error {
	return r.db.Create(content).Error
}

func (r *contentRepository) Update(content *models.Content) error {
	return r.db.Save(content).Error
}

func (r *contentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Content{}, "content_id = ?", id).Error
}
