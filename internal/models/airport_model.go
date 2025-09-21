package models

type Airport struct {
	AirportID uint   `gorm:"primaryKey;autoIncrement;column:airport_id"`
	Name      string `gorm:"size:150;not null;column:name"`
	Code      string `gorm:"size:10;unique;not null;column:code"`
	Address   string `gorm:"size:255:not null;column:address"`
	Timezone  string `gorm:"type:enum('WIB','WITA','WIT');not null;default:'WIB';column:timezone"`

	Users     []*User     `gorm:"foreignKey:AirportID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	Devices   []*Device   `gorm:"foreignKey:AirportID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Contents  []*Content  `gorm:"foreignKey:AirportID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Playlists []*Playlist `gorm:"foreignKey:AirportID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Schedules []*Schedule `gorm:"foreignKey:AirportID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
