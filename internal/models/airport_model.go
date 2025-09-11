package models

type Airport struct {
	AirportID uint   `gorm:"primaryKey;autoIncrement;column:airport_id"`
	Name      string `gorm:"size:150;not null;column:name"`
	Code      string `gorm:"size:10;unique;not null;column:code"`
	Address   string `gorm:"size:255;column:address"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at"`

	Users     []*User     `gorm:"foreignKey:AirportID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	Devices   []*Device   `gorm:"foreignKey:AirportID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Playlists []*Playlist `gorm:"foreignKey:AirportID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
