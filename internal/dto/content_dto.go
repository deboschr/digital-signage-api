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
	Playlists []*SummaryPlaylistDTO `json:"playlists"`
	Airport SummaryAirportDTO `json:"airport"`
}


// ==========================
// Create (Request & Response)
// ==========================

type CreateContentResDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Duration  int    `json:"duration"`
	Airport SummaryAirportDTO `json:"airport"`
}

// ==========================
// Update (Request & Response)
// ==========================

type UpdateContentResDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Duration  int    `json:"duration"`
	Airport SummaryAirportDTO `json:"airport"`
}
