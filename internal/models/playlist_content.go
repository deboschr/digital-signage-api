package models

type PlaylistContent struct {
	PlaylistID uint `gorm:"primaryKey;column:playlist_id"`
	ContentID  uint `gorm:"primaryKey;column:content_id"`
	Order      int  `gorm:"column:order"` // urutan konten dalam playlist

	Playlist Playlist `gorm:"foreignKey:PlaylistID;constraint:OnDelete:CASCADE"`
	Content  Content  `gorm:"foreignKey:ContentID;constraint:OnDelete:CASCADE"`
}
