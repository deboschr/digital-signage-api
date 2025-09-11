package repositories

import (
	"digital_signage_api/internal/models"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	FindAll() ([]models.Schedule, error)
	FindByID(id uint) (*models.Schedule, error)
	Create(schedule *models.Schedule) error
	Update(schedule *models.Schedule) error
	Delete(id uint) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db}
}

func (r *scheduleRepository) FindAll() ([]models.Schedule, error) {
	var schedules []models.Schedule
	err := r.db.Preload("Playlist").Find(&schedules).Error
	return schedules, err
}

func (r *scheduleRepository) FindByID(id uint) (*models.Schedule, error) {
	var schedule models.Schedule
	err := r.db.Preload("Playlist").First(&schedule, "schedule_id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) Create(schedule *models.Schedule) error {
	return r.db.Create(schedule).Error
}

func (r *scheduleRepository) Update(schedule *models.Schedule) error {
	return r.db.Save(schedule).Error
}

func (r *scheduleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Schedule{}, "schedule_id = ?", id).Error
}
