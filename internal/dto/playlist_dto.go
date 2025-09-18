package dto

type GetSummaryPlaylistResDTO struct {
	PlaylistID  uint   `json:"playlist_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GetDetailPlaylistResDTO struct {
	PlaylistID  uint                    `json:"playlist_id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	Airport     GetSummaryAirportResDTO `json:"airport"`
	Contents    []struct {
		GetSummaryContentResDTO
		Order int `json:"order"`
	} `json:"contents"`
}

type CreatePlaylistReqDTO struct {
	AirportID   uint   `json:"airport_id" binding:"required,gt=0"`
	Name        string `json:"name" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"omitempty,max=255"`
}

type UpdatePlaylistReqDTO struct {
	PlaylistID  uint    `json:"playlist_id" binding:"required,gt=0"`
	AirportID   *uint   `json:"airport_id" binding:"omitempty,gt=0"`
	Name        *string `json:"name" binding:"omitempty,min=3,max=100"`
	Description *string `json:"description" binding:"omitempty,max=255"`
}

type CreatePlaylistContentReqDTO struct {
	PlaylistID uint `json:"playlist_id"`
	Contents   []struct {
		ContentID uint `json:"content_id"`
		Order     int  `json:"order"`
	} `json:"contents"`
}

type UpdatePlaylistContentReqDTO struct {
	PlaylistID uint `json:"playlist_id"`
	Contents   []struct {
		ContentID uint `json:"content_id"`
		Order     int  `json:"order"`
	} `json:"contents"`
}

type DeletePlaylistContentReqDTO struct {
	PlaylistID uint   `json:"playlist_id"`
	ContentIDs []uint `json:"content_ids"`
}
