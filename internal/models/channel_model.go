package models

type Channel struct {
	ChannelID   uint   `gorm:"primaryKey;column:channel_id"`
	Name        string `gorm:"size:100;not null;column:name"`
	Description string `gorm:"size:255;column:description"`
	AirportID   uint   `gorm:"not null;column:airport_id"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli;column:updated_at"`
}
