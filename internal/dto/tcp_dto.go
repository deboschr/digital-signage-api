package dto

type PlaylistPayloadDTO struct {
	ScheduleID uint                    `json:"schedule_id"`
	PlaylistID uint                    `json:"playlist_id"`
	Name       string                  `json:"name"`
	Contents   []TCPContentResponseDTO `json:"contents"`
}

type TCPContentResponseDTO struct {
	ContentID uint   `json:"content_id"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	Order     int    `json:"order"`
}
