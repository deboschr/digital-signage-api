package dto

type GetSummaryPlaylistResDTO struct {
	PlaylistID  uint    `json:"playlist_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetDetailPlaylistResDTO struct {
	PlaylistID  uint                        `json:"playlist_id"`
	Name        string                      `json:"name"`
	Description *string                     `json:"description"`
	Airport     GetSummaryAirportResDTO     `json:"airport"`
	Contents    *[]GetPlaylistContentResDTO `json:"contents,omitempty"`
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
	PlaylistID uint `json:"playlist_id" binding:"required,gt=0"`
	Contents   []struct {
		ContentID uint `json:"content_id" binding:"required,gt=0"`
		Order     int  `json:"order" binding:"required,gte=1"`
	} `json:"contents" binding:"required,dive,required"`
}

type UpdatePlaylistContentReqDTO struct {
	PlaylistID uint `json:"playlist_id" binding:"required,gt=0"`
	Contents   []struct {
		ContentID uint `json:"content_id" binding:"required,gt=0"`
		Order     int  `json:"order" binding:"required,gte=1"`
	} `json:"contents" binding:"required,dive,required"`
}

type DeletePlaylistContentReqDTO struct {
	PlaylistID uint   `json:"playlist_id" binding:"required,gt=0"`
	ContentIDs []uint `json:"content_ids" binding:"required,min=1,dive,gt=0"`
}

type GetPlaylistContentResDTO struct {
	GetSummaryContentResDTO
	Order int `json:"order"`
}
