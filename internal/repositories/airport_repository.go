package repositories

import (
	"digital_signage_api/internal/models"

	"gorm.io/gorm"
)

type AirportRepository interface {
	FindAll() ([]*models.Airport, error)
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

func (r *airportRepository) FindAll() ([]*models.Airport, error) {

	var airports []*models.Airport

	err := r.db.Find(&airports).Error

	return airports, err
}

func (r *airportRepository) FindByID(id uint) (*models.Airport, error) {
	var airport models.Airport

	err := r.db.
		Preload("Devices").
		Preload("Playlists").
		Preload("Users").
		Preload("Contents").
		Preload("Schedules").
		First(&airport, "airport_id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &airport, nil
}

func (r *airportRepository) Create(airport *models.Airport) error {
	return r.db.Create(airport).Error
}

func (r *airportRepository) Update(airport *models.Airport) error {
	result := r.db.Save(airport)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (r *airportRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Airport{}, "airport_id = ?", id)
	
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
