package models

import (
	"gorm.io/datatypes"
)

type Schedule struct {
	ScheduleID uint				`gorm:"primaryKey;autoIncrement;column:schedule_id"`
	PlaylistID uint				`gorm:"not null;column:playlist_id"`
	AirportID  uint				`gorm:"not null;column:airport_id"`
	StartTime  int64				`gorm:"not null;column:start_time"`
	EndTime	  int64				`gorm:"not null;column:end_time"`
	RepeatDays datatypes.JSON	`gorm:"type:json;not null;column:repeat_days"`
	IsUrgent   bool				`gorm:"type:tinyint(1);not null;default:0;column:is_urgent"`

	Airport  Airport  `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Playlist Playlist `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
