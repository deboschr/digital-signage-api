package models

type PlaylistContent struct {
	PlaylistID uint `gorm:"primaryKey;column:playlist_id"`
	ContentID  uint `gorm:"primaryKey;column:content_id"`
	Order      int  `gorm:"column:order"`

	Playlist *Playlist `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Content  *Content  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
