package models

type Schedule struct {
	ScheduleID    uint   `gorm:"primaryKey;autoIncrement;column:schedule_id"`
	PlaylistID    uint   `gorm:"not null;column:playlist_id"`
	AirportID     uint   `gorm:"not null;column:airport_id"`
	StartDate     int64  `gorm:"not null;column:start_date"`
	EndDate       int64  `gorm:"not null;column:end_date"`
	StartTime     string `gorm:"type:varchar(8);not null;column:start_time"` // "HH:MM:SS"
	EndTime       string `gorm:"type:varchar(8);not null;column:end_time"`
	RepeatPattern string `gorm:"type:enum('once','daily','weekly','mon','tue','wed','thu','fri','sat','sun');not null;column:repeat_pattern"`
	IsUrgent      bool   `gorm:"type:tinyint(1);not null;default:0;column:is_urgent"`

	Airport  Airport  `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Playlist Playlist `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
