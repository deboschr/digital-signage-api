package models

type Schedule struct {
	ScheduleID uint `gorm:"primaryKey;autoIncrement;column:schedule_id"`
	PlaylistID uint `gorm:"not null;column:playlist_id"`
	AirportID  uint `gorm:"not null;column:airport_id"`

	// Periode aktif (epoch ms)
	StartDate int64 `gorm:"not null;column:start_date"` // kapan mulai berlaku
	EndDate   int64 `gorm:"not null;column:end_date"`   // kapan berhenti berlaku

	// Jam tayang harian
	StartTime string `gorm:"type:time;not null;column:start_time"` // ex: "21:00:00"
	EndTime   string `gorm:"type:time;not null;column:end_time"`   // ex: "03:00:00"

	// Pola perulangan
	RepeatPattern string `gorm:"type:enum('once','daily','weekly','mon','tue','wed','thu','fri','sat','sun');not null;column:repeat_pattern"`
	IsUrgent      bool   `gorm:"type:tinyint(1);not null;default:0;column:is_urgent"`

	Airport  Airport  `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Playlist Playlist `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
