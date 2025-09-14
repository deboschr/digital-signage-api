package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/utils"
)


type ContentService interface {
	GetAllContents() ([]dto.SummaryContentDTO, error)
	GetContentByID(id uint) (dto.DetailContentDTO, error)
	CreateContent(req dto.CreateContentReqDTO) (dto.CreateContentResDTO, error)
	UpdateContent(req dto.UpdateContentReqDTO) (dto.UpdateContentResDTO, error)
	DeleteContent(id uint) error
}

type contentService struct {
	repo repositories.ContentRepository
}

func NewContentService(repo repositories.ContentRepository) ContentService {
	return &contentService{repo}
}

// GET all → Summary DTO
func (s *contentService) GetAllContents() ([]dto.SummaryContentDTO, error) {
	contents, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SummaryContentDTO
	for _, c := range contents {
		res = append(res, dto.SummaryContentDTO{
			ContentID: c.ContentID,
			Title:     c.Title,
			Type:      c.Type,
			Duration:  c.Duration,
			URL:       utils.BuildContentURL(c.Title),
		})
	}
	return res, nil
}

// GET by ID → Detail DTO
func (s *contentService) GetContentByID(id uint) (dto.DetailContentDTO, error) {
	content, err := s.repo.FindByID(id)
	if err != nil {
		return dto.DetailContentDTO{}, err
	}

	playlists := []dto.SummaryPlaylistDTO{}
	for _, p := range content.Playlists {
		playlists = append(playlists, dto.SummaryPlaylistDTO{
			PlaylistID:  p.PlaylistID,
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return dto.DetailContentDTO{
		ContentID: content.ContentID,
		Title:     content.Title,
		Type:      content.Type,
		Duration:  content.Duration,
		URL:       utils.BuildContentURL(content.Title),
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
		Playlists: playlists,
	}, nil
}


// POST → Create DTO
func (s *contentService) CreateContent(req dto.CreateContentReqDTO) (dto.CreateContentResDTO, error) {
	content := models.Content{
		Title:    req.Title,
		Type:     req.Type,
		Duration: req.Duration,
	}

	if err := s.repo.Create(&content); err != nil {
		return dto.CreateContentResDTO{}, err
	}

	return dto.CreateContentResDTO{
		ContentID: content.ContentID,
		Title:     content.Title,
		Type:      content.Type,
		Duration:  content.Duration,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
	}, nil
}

// PUT/PATCH → Update DTO
func (s *contentService) UpdateContent(req dto.UpdateContentReqDTO) (dto.UpdateContentResDTO, error) {
	content, err := s.repo.FindByID(req.ContentID)
	if err != nil {
		return dto.UpdateContentResDTO{}, err
	}

	if req.Title != nil {
		content.Title = *req.Title
	}
	if req.Type != nil {
		content.Type = *req.Type
	}
	if req.Duration != nil {
		content.Duration = *req.Duration
	}

	if err := s.repo.Update(content); err != nil {
		return dto.UpdateContentResDTO{}, err
	}

	return dto.UpdateContentResDTO{
		ContentID: content.ContentID,
		Title:     content.Title,
		Type:      content.Type,
		Duration:  content.Duration,
		CreatedAt: content.CreatedAt,
		UpdatedAt: content.UpdatedAt,
	}, nil
}

// DELETE
func (s *contentService) DeleteContent(id uint) error {
	return s.repo.Delete(id)
}
