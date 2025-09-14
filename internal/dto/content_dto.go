package dto

// ==========================
// Summary untuk listing
// ==========================
type SummaryContentDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Duration  int    `json:"duration"`
	URL       string `json:"url"`
}

// ==========================
// Detail untuk GET /contents/:id
// ==========================
type DetailContentDTO struct {
	ContentID uint                 `json:"content_id"`
	Title     string               `json:"title"`
	Type      string               `json:"type"`
	Duration  int                  `json:"duration"`
	URL       string               `json:"url"`
	CreatedAt int64                `json:"created_at"`
	UpdatedAt int64                `json:"updated_at"`
	Playlists []SummaryPlaylistDTO `json:"playlists"`
}


// ==========================
// Create (Request & Response)
// ==========================
type CreateContentReqDTO struct {
	Title    string `json:"title" binding:"required"`
	Type     string `json:"type" binding:"required,oneof=image video text"`
	Duration int    `json:"duration"`
}

type CreateContentResDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Duration  int    `json:"duration"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ==========================
// Update (Request & Response)
// ==========================
type UpdateContentReqDTO struct {
	ContentID uint    `json:"content_id" binding:"required"`
	Title     *string `json:"title,omitempty"`
	Type      *string `json:"type,omitempty"`
	Duration  *int    `json:"duration,omitempty"`
}


type UpdateContentResDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Duration  int    `json:"duration"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
