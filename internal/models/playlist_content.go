package models

type PlaylistContent struct {
	PlaylistID uint `gorm:"primaryKey;column:playlist_id"`
	ContentID  uint `gorm:"primaryKey;column:content_id"`
	Order      int  `gorm:"column:order"`

	// ON DELETE CASCADE agar pivot ikut hilang kalau parent dihapus
	Playlist Playlist `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Content  Content  `gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}
