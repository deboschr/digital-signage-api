package models

type Playlist struct {
	PlaylistID  uint   `gorm:"primaryKey;autoIncrement;column:playlist_id"`
	AirportID   uint   `gorm:"not null;column:airport_id"`
	Name        string `gorm:"size:100;not null;column:name"`
	Description string `gorm:"size:255;column:description"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli;column:updated_at"`

	Airport Airport `gorm:"foreignKey:AirportID;references:AirportID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
