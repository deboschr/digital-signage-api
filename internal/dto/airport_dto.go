package dto

type GetSummaryAirportResDTO struct {
	AirportID uint   `json:"airport_id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Address   string `json:"address"`
	Timezone  string `json:"timezone"`
}

type GetDetailAirportResDTO struct {
	AirportID uint                        `json:"airport_id"`
	Name      string                      `json:"name"`
	Code      string                      `json:"code"`
	Address   string                      `json:"address"`
	Timezone  string                      `json:"timezone"`
	Users     []*GetSummaryUserResDTO     `json:"users"`
	Devices   []*GetSummaryDeviceResDTO   `json:"devices"`
	Contents  []*GetSummaryContentResDTO  `json:"contents"`
	Playlists []*GetSummaryPlaylistResDTO `json:"playlists"`
	Schedules []*GetSummaryScheduleResDTO `json:"schedules"`
}

type CreateAirportReqDTO struct {
	Name     string `json:"name" binding:"required,min=3,max=150"`
	Code     string `json:"code" binding:"required,alphanum,len=3"`
	Address  string `json:"address" binding:"required,max=255"`
	Timezone string `json:"timezone" binding:"required,oneof=WIB WITA WIT"`
}

type UpdateAirportReqDTO struct {
	AirportID uint    `json:"airport_id" binding:"required,gt=0"`
	Name      *string `json:"name" binding:"omitempty,min=3,max=150"`
	Code      *string `json:"code" binding:"omitempty,alphanum,len=3"`
	Address   *string `json:"address" binding:"omitempty,max=255"`
	Timezone  *string `json:"timezone" binding:"omitempty,oneof=WIB WITA WIT"`
}
