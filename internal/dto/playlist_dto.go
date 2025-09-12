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
// Detail untuk GET /playlists/:id
// ==========================
type DetailPlaylistDTO struct {
	PlaylistID  uint                  `json:"playlist_id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	CreatedAt   int64                 `json:"created_at"`
	UpdatedAt   int64                 `json:"updated_at"`
	Airport     *SummaryAirportDTO    `json:"airport,omitempty"`
	Contents    []SummaryContentDTO   `json:"contents"`
	Schedules   []SummaryScheduleDTO  `json:"schedules"`
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
