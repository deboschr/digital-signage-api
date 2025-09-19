package models

type Playlist struct {
	PlaylistID  uint    `gorm:"primaryKey;autoIncrement;column:playlist_id"`
	AirportID   uint    `gorm:"not null;column:airport_id"`
	Name        string  `gorm:"size:100;not null;column:name"`
	Description *string `gorm:"size:255;column:description"`

	Airport         Airport            `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Contents        []*Content         `gorm:"many2many:playlist_contents;joinForeignKey:PlaylistID;joinReferences:ContentID"`
	Schedules       []*Schedule        `gorm:"foreignKey:PlaylistID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
