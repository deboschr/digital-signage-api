package services

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type ScheduleService interface {
	GetAllSchedules() ([]models.Schedule, error)
	GetScheduleByID(id uint) (*models.Schedule, error)
	CreateSchedule(schedule *models.Schedule) error
	UpdateSchedule(schedule *models.Schedule) error
	DeleteSchedule(id uint) error
}

type scheduleService struct {
	repo repositories.ScheduleRepository
}

func NewScheduleService(repo repositories.ScheduleRepository) ScheduleService {
	return &scheduleService{repo}
}

func (s *scheduleService) GetAllSchedules() ([]models.Schedule, error) {
	return s.repo.FindAll()
}

func (s *scheduleService) GetScheduleByID(id uint) (*models.Schedule, error) {
	return s.repo.FindByID(id)
}

func (s *scheduleService) CreateSchedule(schedule *models.Schedule) error {
	return s.repo.Create(schedule)
}

func (s *scheduleService) UpdateSchedule(schedule *models.Schedule) error {
	return s.repo.Update(schedule)
}

func (s *scheduleService) DeleteSchedule(id uint) error {
	return s.repo.Delete(id)
}
