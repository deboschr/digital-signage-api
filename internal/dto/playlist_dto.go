package dto

// ==========================
// Summary untuk listing
// ==========================
type SummaryPlaylistDTO struct {
	PlaylistID  uint   `json:"playlist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ==========================
// Item khusus untuk konten dalam playlist
// ==========================
type PlaylistContentItemDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	Order     int    `json:"order"`
	Duration  int    `json:"duration"`
}

// ==========================
// Detail untuk GET /playlists/:id
// ==========================
type DetailPlaylistDTO struct {
	PlaylistID  uint                    `json:"playlist_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	CreatedAt   int64                   `json:"created_at"`
	UpdatedAt   int64                   `json:"updated_at"`
	Airport     *SummaryAirportDTO      `json:"airport,omitempty"`
	Contents    []PlaylistContentItemDTO `json:"contents"`
	Schedules   []SummaryScheduleDTO    `json:"schedules"`
}

// ==========================
// Create (Request & Response)
// ==========================
type CreatePlaylistReqDTO struct {
	AirportID   uint   `json:"airport_id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type CreatePlaylistResDTO struct {
	PlaylistID  uint   `json:"playlist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ==========================
// Update (Request & Response)
// ==========================
type UpdatePlaylistReqDTO struct {
	PlaylistID  uint    `json:"playlist_id" binding:"required"`
	AirportID   *uint   `json:"airport_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdatePlaylistResDTO struct {
	PlaylistID  uint   `json:"playlist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}

// ==========================
// Playlist Content DTO
// ==========================

// POST /playlists/content
type CreatePlaylistContentReqDTO struct {
	PlaylistID uint   `json:"playlist_id" binding:"required"`
	ContentIDs []uint `json:"content_ids" binding:"required"`
}

// PATCH /playlists/content
type UpdatePlaylistContentReqDTO struct {
	PlaylistID uint `json:"playlist_id" binding:"required"`
	Contents   []struct {
		ContentID uint `json:"content_id" binding:"required"`
		Order     int  `json:"order" binding:"required"`
	} `json:"contents" binding:"required"`
}

// DELETE /playlists/content
type DeletePlaylistContentReqDTO struct {
	PlaylistID uint   `json:"playlist_id" binding:"required"`
	ContentIDs []uint `json:"content_ids" binding:"required"`
}
