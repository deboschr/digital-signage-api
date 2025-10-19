package dto

type GetSummaryScheduleResDTO struct {
	ScheduleID    uint     `json:"schedule_id"`
	Playlist      string   `json:"playlist"`
	StartTime     string   `json:"start_time"`
	EndTime       string   `json:"end_time"`
	RepeatPattern []string `json:"repeat_pattern"`
	IsUrgent      bool     `json:"is_urgent"`
}

type GetDetailScheduleResDTO struct {
	ScheduleID    uint                     `json:"schedule_id"`
	StartTime     string                   `json:"start_time"`
	EndTime       string                   `json:"end_time"`
	RepeatPattern []string                 `json:"repeat_pattern"`
	IsUrgent      bool                     `json:"is_urgent"`
	Playlist      GetSummaryPlaylistResDTO `json:"playlist"`
	Airport       GetSummaryAirportResDTO  `json:"airport"`
}

type CreateScheduleReqDTO struct {
	PlaylistID    uint     `json:"playlist_id" binding:"required,gt=0"`
	AirportID     uint     `json:"airport_id" binding:"required,gt=0"`
	StartTime     int64    `json:"start_time" binding:"required,gte=0"`
	EndTime       int64    `json:"end_time" binding:"required,gte=0"`
	RepeatPattern []string `json:"repeat_pattern" binding:"required,dive,oneof=once mon tue wed thu fri sat sun"`
	IsUrgent      bool     `json:"is_urgent" binding:"omitempty"`
}

type UpdateScheduleReqDTO struct {
	ScheduleID    uint      `json:"schedule_id" binding:"required,gt=0"`
	PlaylistID    *uint     `json:"playlist_id" binding:"omitempty,gt=0"`
	AirportID     *uint     `json:"airport_id" binding:"required,gt=0"`
	StartTime     *int64    `json:"start_time" binding:"omitempty,gte=0"`
	EndTime       *int64    `json:"end_time" binding:"omitempty,gte=0"`
	RepeatPattern *[]string `json:"repeat_pattern" binding:"omitempty,dive,oneof=once mon tue wed thu fri sat sun"`
	IsUrgent      *bool     `json:"is_urgent" binding:"omitempty"`
}

type ActiveScheduleRes struct {
	ScheduleID    uint     `json:"schedule_id"`
	PlaylistID    uint     `json:"playlist_id"`
	Name          string   `json:"name"`
	IsUrgent      bool     `json:"is_urgent"`
	StartTime     string   `json:"start_time"`
	EndTime       string   `json:"end_time"`
	Contents      []struct {
		ContentID uint   `json:"content_id"`
		Title     string `json:"title"`
		URL       string `json:"url"`
		Order     int    `json:"order"`
	} `json:"contents"`
}
