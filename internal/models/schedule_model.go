package models

type Schedule struct {
	ScheduleID    uint   `gorm:"primaryKey;column:schedule_id"`
	ChannelID     uint   `gorm:"not null;column:channel_id"`
	ContentID     uint   `gorm:"not null;column:content_id"`
	StartTime     int64  `gorm:"not null;column:start_time"`    // epoch ms
	EndTime       int64  `gorm:"not null;column:end_time"`      // epoch ms
	RepeatPattern string `gorm:"size:50;column:repeat_pattern"` // daily, once, weekly
	CreatedAt     int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt     int64  `gorm:"autoUpdateTime:milli;column:updated_at"`
}