package models

type Playlist struct {
	PlaylistID  uint    `gorm:"primaryKey;autoIncrement;column:playlist_id"`
	AirportID   uint    `gorm:"not null;column:airport_id"`
	Name        string  `gorm:"size:100;not null;column:name"`
	Description *string `gorm:"size:255;column:description"`

	Airport          Airport            `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	PlaylistContents []*PlaylistContent `gorm:"foreignKey:PlaylistID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Schedules        []*Schedule        `gorm:"foreignKey:PlaylistID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
