package models

type Content struct {
	ContentID uint   `gorm:"primaryKey;autoIncrement;column:content_id"`
	AirportID uint   `gorm:"not null;column:airport_id"`
	Title     string `gorm:"size:150;not null;column:title"`
	Type      string `gorm:"type:enum('image','video');not null;column:type"`
	Duration  uint16 `gorm:"column:duration;not null"` // detik, 0–3600

	Airport          Airport            `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	PlaylistContents []*PlaylistContent `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
