package models

type Content struct {
	ContentID uint   `gorm:"primaryKey;column:content_id"`
	Title     string `gorm:"size:150;not null;column:title"`
	Type      string `gorm:"type:enum('image','video','text');not null;column:type"`
	FileURL   string `gorm:"size:255;not null;column:file_url"`
	Duration  int    `gorm:"column:duration"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at"`

	Playlists []Playlist `gorm:"many2many:playlist_contents;joinForeignKey:ContentID;JoinReferences:PlaylistID"`
}
