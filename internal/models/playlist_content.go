package models

type PlaylistContent struct {
	PlaylistID uint `gorm:"primaryKey;column:playlist_id"`
	ContentID  uint `gorm:"primaryKey;column:content_id"`
	Order      int  `gorm:"column:order"`

	Playlist Playlist `gorm:"foreignKey:PlaylistID;references:PlaylistID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Content  Content  `gorm:"foreignKey:ContentID;references:ContentID;constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
