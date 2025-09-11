package services

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type AirportService interface {
	GetAllAirports() ([]models.Airport, error)
	GetAirportByID(id uint) (*models.Airport, error)
	CreateAirport(airport *models.Airport) error
	UpdateAirport(airport *models.Airport) error
	DeleteAirport(id uint) error
}

type airportService struct {
	repo repositories.AirportRepository
}

func NewAirportService(repo repositories.AirportRepository) AirportService {
	return &airportService{repo}
}

func (s *airportService) GetAllAirports() ([]models.Airport, error) {
	return s.repo.FindAll()
}

func (s *airportService) GetAirportByID(id uint) (*models.Airport, error) {
	return s.repo.FindByID(id)
}

func (s *airportService) CreateAirport(airport *models.Airport) error {
	return s.repo.Create(airport)
}

func (s *airportService) UpdateAirport(airport *models.Airport) error {
	return s.repo.Update(airport)
}

func (s *airportService) DeleteAirport(id uint) error {
	return s.repo.Delete(id)
}
