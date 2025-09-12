package repositories

import (
	"digital_signage_api/internal/models"

	"gorm.io/gorm"
)

type AirportRepository interface {
	FindAll() ([]models.Airport, error)
	FindByID(id uint) (*models.Airport, error)
	Create(airport *models.Airport) error
	Update(airport *models.Airport) error
	Delete(id uint) error
}

type airportRepository struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) AirportRepository {
	return &airportRepository{db}
}

// Untuk summary DTO → cukup field dasar
func (r *airportRepository) FindAll() ([]models.Airport, error) {
	var airports []models.Airport
	err := r.db.
		Select("airport_id", "name", "code", "address").
		Find(&airports).Error
	return airports, err
}

// Untuk detail DTO → preload relasi penuh
func (r *airportRepository) FindByID(id uint) (*models.Airport, error) {
	var airport models.Airport
	err := r.db.
		Preload("Devices").
		Preload("Playlists").
		Preload("Users").
		First(&airport, "airport_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &airport, nil
}

func (r *airportRepository) Create(airport *models.Airport) error {
	return r.db.Create(airport).Error
}

func (r *airportRepository) Update(airport *models.Airport) error {
	return r.db.Save(airport).Error
}

func (r *airportRepository) Delete(id uint) error {
	return r.db.Delete(&models.Airport{}, "airport_id = ?", id).Error
}
