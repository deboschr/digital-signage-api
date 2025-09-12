package dto

// ==========================
// Summary untuk listing
// ==========================
type SummaryScheduleDTO struct {
	ScheduleID    uint   `json:"schedule_id"`
	StartTime     int64  `json:"start_time"`
	EndTime       int64  `json:"end_time"`
	RepeatPattern string `json:"repeat_pattern"`
}

// ==========================
// Detail untuk GET /schedules/:id
// ==========================
type DetailScheduleDTO struct {
	ScheduleID    uint               `json:"schedule_id"`
	StartTime     int64              `json:"start_time"`
	EndTime       int64              `json:"end_time"`
	RepeatPattern string             `json:"repeat_pattern"`
	CreatedAt     int64              `json:"created_at"`
	UpdatedAt     int64              `json:"updated_at"`
	Playlist      *SummaryPlaylistDTO `json:"playlist,omitempty"`
}

// ==========================
// Create (Request & Response)
// ==========================
type CreateScheduleReqDTO struct {
	PlaylistID    uint   `json:"playlist_id" binding:"required"`
	StartTime     int64  `json:"start_time" binding:"required"`
	EndTime       int64  `json:"end_time" binding:"required"`
	RepeatPattern string `json:"repeat_pattern"`
}

type CreateScheduleResDTO struct {
	ScheduleID    uint   `json:"schedule_id"`
	StartTime     int64  `json:"start_time"`
	EndTime       int64  `json:"end_time"`
	RepeatPattern string `json:"repeat_pattern"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}

// ==========================
// Update (Request & Response)
// ==========================
type UpdateScheduleReqDTO struct {
	ScheduleID    uint    `json:"schedule_id" binding:"required"`
	PlaylistID    *uint   `json:"playlist_id,omitempty"`
	StartTime     *int64  `json:"start_time,omitempty"`
	EndTime       *int64  `json:"end_time,omitempty"`
	RepeatPattern *string `json:"repeat_pattern,omitempty"`
}


type UpdateScheduleResDTO struct {
	ScheduleID    uint   `json:"schedule_id"`
	StartTime     int64  `json:"start_time"`
	EndTime       int64  `json:"end_time"`
	RepeatPattern string `json:"repeat_pattern"`
	CreatedAt     int64  `json:"created_at"`
	UpdatedAt     int64  `json:"updated_at"`
}
