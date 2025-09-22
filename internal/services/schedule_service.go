package services

import (
	"digital_signage_api/internal/dto"
	"digital_signage_api/internal/models"
	"digital_signage_api/internal/repositories"
	// "time"
)

type ScheduleService interface {
	GetSchedules() ([]dto.GetSummaryScheduleResDTO, error)
	GetSchedule(id uint) (dto.GetDetailScheduleResDTO, error)
	CreateSchedule(req dto.CreateScheduleReqDTO) (dto.GetSummaryScheduleResDTO, error)
	UpdateSchedule(req dto.UpdateScheduleReqDTO) (dto.GetSummaryScheduleResDTO, error)
	DeleteSchedule(id uint) error
}

type scheduleService struct {
	repo repositories.ScheduleRepository
}

func NewScheduleService(repo repositories.ScheduleRepository) ScheduleService {
	return &scheduleService{repo}
}

func (s *scheduleService) GetSchedules() ([]dto.GetSummaryScheduleResDTO, error) {

	schedules, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var res []dto.GetSummaryScheduleResDTO
	for _, sch := range schedules {
		res = append(res, dto.GetSummaryScheduleResDTO{
			ScheduleID:    sch.ScheduleID,
			Playlist:     sch.Playlist.Name,
			StartTime:     sch.StartTime,
			EndTime:       sch.EndTime,
			RepeatPattern: sch.RepeatPattern,
			IsUrgent: sch.IsUrgent,
		})
	}

	return res, nil
}

func (s *scheduleService) GetSchedule(id uint) (dto.GetDetailScheduleResDTO, error) {

	schedule, err := s.repo.FindByID(id)

	if err != nil {
		return dto.GetDetailScheduleResDTO{}, err
	}

	playlist := dto.GetSummaryPlaylistResDTO{
		PlaylistID:  schedule.Playlist.PlaylistID,
		Name:        schedule.Playlist.Name,
		Description: schedule.Playlist.Description,
	}

	    // Mapping airport
   airport := dto.GetSummaryAirportResDTO{
      AirportID: schedule.Airport.AirportID,
   	Name:      schedule.Airport.Name,
      Code:      schedule.Airport.Code,
      Address:   schedule.Airport.Address,
   }

	return dto.GetDetailScheduleResDTO{
		ScheduleID:    schedule.ScheduleID,
		StartTime: schedule.StartTime,
		EndTime:   schedule.EndTime,
		RepeatPattern: schedule.RepeatPattern,
		IsUrgent: schedule.IsUrgent,
		Playlist:      playlist,
		Airport:      airport,
	}, nil
}

func (s *scheduleService) CreateSchedule(req dto.CreateScheduleReqDTO) (dto.GetSummaryScheduleResDTO, error) {
	
	// layout := "15:04:05"
	// startTime, _ := time.Parse(layout, req.StartTime)
	// endTime, _ := time.Parse(layout, req.EndTime)

	schedule := models.Schedule{
		PlaylistID:    req.PlaylistID,
		AirportID:     req.AirportID,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
    	StartTime:     req.StartTime,
    	EndTime:       req.EndTime,
		RepeatPattern: req.RepeatPattern,
		IsUrgent:      req.IsUrgent,
	}

	if err := s.repo.Create(&schedule); err != nil {
		return dto.GetSummaryScheduleResDTO{}, err
	}

	return dto.GetSummaryScheduleResDTO{
		ScheduleID:    schedule.ScheduleID,
		StartDate:     schedule.StartDate,
		EndDate:       schedule.EndDate,
		StartTime:     schedule.StartTime,
		EndTime:       schedule.EndTime,
		RepeatPattern: schedule.RepeatPattern,
		IsUrgent:      schedule.IsUrgent,
	}, nil
}

func (s *scheduleService) UpdateSchedule(req dto.UpdateScheduleReqDTO) (dto.GetSummaryScheduleResDTO, error) {
	
	schedule, err := s.repo.FindByID(req.ScheduleID)
	
	if err != nil {
		return dto.GetSummaryScheduleResDTO{}, err
	}

	if req.AirportID != nil {
		schedule.AirportID = *req.AirportID
	}
	if req.PlaylistID != nil {
		schedule.PlaylistID = *req.PlaylistID
	}
	if req.StartDate != nil {
		schedule.StartDate = *req.StartDate
	}
	if req.EndDate != nil {
		schedule.EndDate = *req.EndDate
	}
	if req.StartTime != nil {
    	// layout := "15:04:05"
    	// startTime, _ := time.Parse(layout, *req.StartTime)
    	schedule.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
    	// layout := "15:04:05"
    	// endTime, _ := time.Parse(layout, *req.EndTime)
    	schedule.EndTime = *req.EndTime
	}

	if req.RepeatPattern != nil {
		schedule.RepeatPattern = *req.RepeatPattern
	}
	if req.IsUrgent != nil {
		schedule.IsUrgent = *req.IsUrgent
	}

	if err := s.repo.Update(schedule); err != nil {
		return dto.GetSummaryScheduleResDTO{}, err
	}

	return dto.GetSummaryScheduleResDTO{
		ScheduleID:    schedule.ScheduleID,
		StartDate:     schedule.StartDate,
		EndDate:       schedule.EndDate,
		StartTime:     schedule.StartTime,
		EndTime:       schedule.EndTime,
		RepeatPattern: schedule.RepeatPattern,
		IsUrgent:      schedule.IsUrgent,
	}, nil
}

func (s *scheduleService) DeleteSchedule(id uint) error {
	return s.repo.Delete(id)
}
