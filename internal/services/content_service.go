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

func (s *contentService) GetContent(id uint) (dto.GetDetailContentResDTO, error) {
    content, err := s.repo.FindByID(id)
    if err != nil {
        return dto.GetDetailContentResDTO{}, err
    }

    // Mapping playlists
    playlists := []*dto.GetSummaryPlaylistResDTO{}
    for _, p := range content.Playlists {
        playlist := dto.GetSummaryPlaylistResDTO{
            PlaylistID:  p.PlaylistID,
            Name:        p.Name,
            Description: p.Description,
        }
        playlists = append(playlists, &playlist)
    }

    // Mapping airport
    airport := dto.GetSummaryAirportResDTO{
        AirportID: content.Airport.AirportID,
        Name:      content.Airport.Name,
        Code:      content.Airport.Code,
        Address:   content.Airport.Address,
    }

    return dto.GetDetailContentResDTO{
        ContentID: content.ContentID,
        Title:     content.Title,
        Type:      content.Type,
        Duration:  content.Duration,
        URL:       utils.BuildContentURL(content.Title),
        Playlists: playlists,
        Airport:   airport,
    }, nil
}



func (s *contentService) CreateContent(req dto.CreateContentReqDTO) (dto.GetSummaryContentResDTO, error) {
	content := models.Content{
		AirportID:    req.AirportID,
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


func (s *contentService) DeleteContent(id uint) error {
	return s.repo.Delete(id)
}
