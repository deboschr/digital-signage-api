package dto

type GetSummaryAirportResDTO struct {
	AirportID uint    `json:"airport_id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Address   *string `json:"address"`
}

type GetDetailAirportResDTO struct {
	AirportID uint                        `json:"airport_id"`
	Name      string                      `json:"name"`
	Code      string                      `json:"code"`
	Address   *string                     `json:"address"`
	Users     []*GetSummaryUserResDTO     `json:"users"`
	Devices   []*GetSummaryDeviceResDTO   `json:"devices"`
	Contents  []*GetSummaryContentResDTO  `json:"contents"`
	Playlists []*GetSummaryPlaylistResDTO `json:"playlists"`
	Schedules []*GetSummaryScheduleResDTO `json:"schedules"`
}

type CreateAirportReqDTO struct {
	Name    string  `json:"name" binding:"required,min=3,max=150"`
	Code    string  `json:"code" binding:"required,alphanum,len=3"`
	Address *string `json:"address" binding:"omitempty,max=255"`
}

type UpdateAirportReqDTO struct {
	AirportID uint    `json:"airport_id" binding:"required,gt=0"`
	Name      *string `json:"name" binding:"omitempty,min=3,max=150"`
	Code      *string `json:"code" binding:"omitempty,alphanum,len=3"`
	Address   *string `json:"address" binding:"omitempty,max=255"`
}
