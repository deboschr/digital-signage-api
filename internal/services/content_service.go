package services

import (
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type ContentService interface {
	GetAllContents() ([]models.Content, error)
	GetContentByID(id uint) (*models.Content, error)
	CreateContent(content *models.Content) error
	UpdateContent(content *models.Content) error
	DeleteContent(id uint) error
}

type contentService struct {
	repo repositories.ContentRepository
}

func NewContentService(repo repositories.ContentRepository) ContentService {
	return &contentService{repo}
}

func (s *contentService) GetAllContents() ([]models.Content, error) {
	return s.repo.FindAll()
}

func (s *contentService) GetContentByID(id uint) (*models.Content, error) {
	return s.repo.FindByID(id)
}

func (s *contentService) CreateContent(content *models.Content) error {
	return s.repo.Create(content)
}

func (s *contentService) UpdateContent(content *models.Content) error {
	return s.repo.Update(content)
}

func (s *contentService) DeleteContent(id uint) error {
	return s.repo.Delete(id)
}
