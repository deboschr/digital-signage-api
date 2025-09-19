package dto

type GetSummaryScheduleResDTO struct {
	ScheduleID    uint   `json:"schedule_id"`
	StartDate     int64  `json:"start_date"`
	EndDate       int64  `json:"end_date"`
	StartTime     string `json:"start_time"`
	EndTime       string `json:"end_time"`
	RepeatPattern string `json:"repeat_pattern"`
	IsUrgent      bool   `json:"is_urgent"`
}

type GetDetailScheduleResDTO struct {
	ScheduleID    uint                     `json:"schedule_id"`
	StartDate     int64                    `json:"start_date"`
	EndDate       int64                    `json:"end_date"`
	StartTime     string                   `json:"start_time"`
	EndTime       string                   `json:"end_time"`
	RepeatPattern string                   `json:"repeat_pattern"`
	IsUrgent      bool                     `json:"is_urgent"`
	Playlist      GetSummaryPlaylistResDTO `json:"playlist"`
	Airport       GetSummaryAirportResDTO  `json:"airport"`
}

type CreateScheduleReqDTO struct {
	PlaylistID    uint   `json:"playlist_id" binding:"required,gt=0"`
	StartDate     int64  `json:"start_date" binding:"required,gte=0"`
	EndDate       int64  `json:"end_date" binding:"required,gte=0"`
	StartTime     string `json:"start_time" binding:"required,datetime=15:04:05"`
	EndTime       string `json:"end_time" binding:"required,datetime=15:04:05"`
	RepeatPattern string `json:"repeat_pattern" binding:"required,oneof=once daily weekly mon tue wed thu fri sat sun"`
	IsUrgent      *bool  `json:"is_urgent" binding:"omitempty"`
}

type UpdateScheduleReqDTO struct {
	ScheduleID    uint    `json:"schedule_id" binding:"required,gt=0"`
	PlaylistID    *uint   `json:"playlist_id" binding:"omitempty,gt=0"`
	StartDate     *int64  `json:"start_date" binding:"omitempty,gte=0"`
	EndDate       *int64  `json:"end_date" binding:"omitempty,gte=0"`
	StartTime     *string `json:"start_time" binding:"omitempty,datetime=15:04:05"`
	EndTime       *string `json:"end_time" binding:"omitempty,datetime=15:04:05"`
	RepeatPattern *string `json:"repeat_pattern" binding:"omitempty,oneof=once daily weekly mon tue wed thu fri sat sun"`
	IsUrgent      *bool   `json:"is_urgent" binding:"omitempty"`
}

type ActiveScheduleRes struct {
	ScheduleID uint   `json:"schedule_id"`
	PlaylistID uint   `json:"playlist_id"`
	Name       string `json:"name"`
	Contents   []struct {
		ContentID uint   `json:"content_id"`
		Title     string `json:"title"`
		URL       string `json:"url"`
		Order     int    `json:"order"`
	} `json:"contents"`
}
