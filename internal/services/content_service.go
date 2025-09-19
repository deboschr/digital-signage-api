package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
	"digital_signage_api/internal/utils"
)


type ContentService interface {
	GetContents() ([]dto.GetSummaryContentResDTO, error)
	GetContent(id uint) (dto.GetDetailContentResDTO, error)
	CreateContent(req dto.CreateContentReqDTO) (dto.GetSummaryContentResDTO, error)
	DeleteContent(id uint) error
}

type contentService struct {
	repo repositories.ContentRepository
}

func NewContentService(repo repositories.ContentRepository) ContentService {
	return &contentService{repo}
}

func (s *contentService) GetContents() ([]dto.GetSummaryContentResDTO, error) {
	contents, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.GetSummaryContentResDTO
	for _, c := range contents {
		res = append(res, dto.GetSummaryContentResDTO{
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
func (s *contentService) GetContent(id uint) (dto.GetDetailContentResDTO, error) {
	content, err := s.repo.FindByID(id)
	if err != nil {
		return dto.GetDetailContentResDTO{}, err
	}

	playlists := []dto.SummaryPlaylistDTO{}
	for _, p := range content.Playlists {
		playlists = append(playlists, dto.SummaryPlaylistDTO{
			PlaylistID:  p.PlaylistID,
			Name:        p.Name,
			Description: p.Description,
		})
	}

	return dto.GetDetailContentResDTO{
		ContentID: content.ContentID,
		Title:     content.Title,
		Type:      content.Type,
		Duration:  content.Duration,
		URL:       utils.BuildContentURL(content.Title),
		Playlists: playlists,
	}, nil
}


// POST → Create DTO
func (s *contentService) CreateContent(req dto.CreateContentReqDTO) (dto.GetSummaryContentResDTO, error) {
	content := models.Content{
		Title:    req.Title,
		Type:     req.Type,
		Duration: req.Duration,
	}

	if err := s.repo.Create(&content); err != nil {
		return dto.GetSummaryContentResDTO{}, err
	}

	return dto.GetSummaryContentResDTO{
		ContentID: content.ContentID,
		Title:     content.Title,
		Type:      content.Type,
		Duration:  content.Duration,
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
