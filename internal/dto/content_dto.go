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

type CreateContentReqDTO struct {
	AirportID uint   `json:"airport_id" binding:"required,gt=0"`
	Title     string `json:"title" binding:"required,min=3,max=150"`
	Type      string `json:"type" binding:"required,oneof=image video"`
	Duration  uint16 `json:"duration" binding:"required,gte=0,lte=3600"`
}
