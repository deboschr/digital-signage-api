package dto

type GetAirportsDTO struct {
	AirportID uint    `json:"airport_id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Address   *string `json:"address"`
}

type GetAirportDTO struct {
	AirportID uint                  `json:"airport_id"`
	Name      string                `json:"name"`
	Code      string                `json:"code"`
	Address   *string               `json:"address"`
	Users     []*SummaryUserDTO     `json:"users"`
	Devices   []*SummaryDeviceDTO   `json:"devices"`
	Playlists []*SummaryPlaylistDTO `json:"playlists"`
	Contents  []*SummaryContentDTO  `json:"contents"`
}

type CreateAirportDTO struct {
	Name    string  `json:"name" binding:"required"`
	Code    string  `json:"code" binding:"required"`
	Address *string `json:"address"`
}

type UpdateAirportDTO struct {
	AirportID uint    `json:"airport_id" binding:"required"`
	Name      *string `json:"name,omitempty"`
	Code      *string `json:"code,omitempty"`
	Address   *string `json:"address,omitempty"`
}