package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
)

type ScheduleService interface {
	GetAllSchedules() ([]dto.SummaryScheduleDTO, error)
	GetScheduleByID(id uint) (dto.DetailScheduleDTO, error)
	CreateSchedule(req dto.CreateScheduleReqDTO) (dto.CreateScheduleResDTO, error)
	UpdateSchedule(req dto.UpdateScheduleReqDTO) (dto.UpdateScheduleResDTO, error)
	DeleteSchedule(id uint) error
}

type scheduleService struct {
	repo repositories.ScheduleRepository
}

func NewScheduleService(repo repositories.ScheduleRepository) ScheduleService {
	return &scheduleService{repo}
}

// GET all → Summary DTO
func (s *scheduleService) GetAllSchedules() ([]dto.SummaryScheduleDTO, error) {
	schedules, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var res []dto.SummaryScheduleDTO
	for _, sch := range schedules {
		res = append(res, dto.SummaryScheduleDTO{
			ScheduleID:    sch.ScheduleID,
			StartTime:     sch.StartTime,
			EndTime:       sch.EndTime,
			RepeatPattern: sch.RepeatPattern,
		})
	}
	return res, nil
}

// GET by ID → Detail DTO
func (s *scheduleService) GetScheduleByID(id uint) (dto.DetailScheduleDTO, error) {
	schedule, err := s.repo.FindByID(id)
	if err != nil {
		return dto.DetailScheduleDTO{}, err
	}

	var playlist *dto.SummaryPlaylistDTO
	if schedule.Playlist != nil {
		playlist = &dto.SummaryPlaylistDTO{
			PlaylistID:  schedule.Playlist.PlaylistID,
			Name:        schedule.Playlist.Name,
			Description: schedule.Playlist.Description,
		}
	}

	return dto.DetailScheduleDTO{
		ScheduleID:    schedule.ScheduleID,
		StartTime:     schedule.StartTime,
		EndTime:       schedule.EndTime,
		RepeatPattern: schedule.RepeatPattern,
		CreatedAt:     schedule.CreatedAt,
		UpdatedAt:     schedule.UpdatedAt,
		Playlist:      playlist,
	}, nil
}

// POST → Create DTO
func (s *scheduleService) CreateSchedule(req dto.CreateScheduleReqDTO) (dto.CreateScheduleResDTO, error) {
	schedule := models.Schedule{
		PlaylistID:    req.PlaylistID,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		RepeatPattern: req.RepeatPattern,
	}

	if err := s.repo.Create(&schedule); err != nil {
		return dto.CreateScheduleResDTO{}, err
	}

	return dto.CreateScheduleResDTO{
		ScheduleID:    schedule.ScheduleID,
		StartTime:     schedule.StartTime,
		EndTime:       schedule.EndTime,
		RepeatPattern: schedule.RepeatPattern,
		CreatedAt:     schedule.CreatedAt,
		UpdatedAt:     schedule.UpdatedAt,
	}, nil
}

// PUT/PATCH → Update DTO
func (s *scheduleService) UpdateSchedule(req dto.UpdateScheduleReqDTO) (dto.UpdateScheduleResDTO, error) {
	schedule, err := s.repo.FindByID(req.ScheduleID)
	if err != nil {
		return dto.UpdateScheduleResDTO{}, err
	}

	if req.PlaylistID != nil {
		schedule.PlaylistID = *req.PlaylistID
	}
	if req.StartTime != nil {
		schedule.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
		schedule.EndTime = *req.EndTime
	}
	if req.RepeatPattern != nil {
		schedule.RepeatPattern = *req.RepeatPattern
	}

	if err := s.repo.Update(schedule); err != nil {
		return dto.UpdateScheduleResDTO{}, err
	}

	return dto.UpdateScheduleResDTO{
		ScheduleID:    schedule.ScheduleID,
		StartTime:     schedule.StartTime,
		EndTime:       schedule.EndTime,
		RepeatPattern: schedule.RepeatPattern,
		CreatedAt:     schedule.CreatedAt,
		UpdatedAt:     schedule.UpdatedAt,
	}, nil
}

// DELETE
func (s *scheduleService) DeleteSchedule(id uint) error {
	return s.repo.Delete(id)
}
