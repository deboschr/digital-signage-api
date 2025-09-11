package models

type Schedule struct {
	ScheduleID    uint   `gorm:"primaryKey;autoIncrement;column:schedule_id"`
	PlaylistID    uint   `gorm:"not null;column:playlist_id"`
	StartTime     int64  `gorm:"not null;column:start_time"`
	EndTime       int64  `gorm:"not null;column:end_time"`
	RepeatPattern string `gorm:"size:50;column:repeat_pattern"`
	CreatedAt     int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt     int64  `gorm:"autoUpdateTime:milli;column:updated_at"`

	Playlist Playlist `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
