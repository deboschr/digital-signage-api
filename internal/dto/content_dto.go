package dto

type GetSummaryContentResDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Duration  uint16 `json:"duration"`
	URL       string `json:"url"`
}

type GetDetailContentResDTO struct {
	ContentID uint                        `json:"content_id"`
	Title     string                      `json:"title"`
	Type      string                      `json:"type"`
	Duration  int                         `json:"duration"`
	URL       string                      `json:"url"`
	Airport   GetSummaryAirportResDTO     `json:"airport"`
	Playlists []*GetSummaryPlaylistResDTO `json:"playlists"`
}