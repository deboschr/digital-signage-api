package dto

// ==========================
// Summary untuk listing
// ==========================
type SummaryAirportDTO struct {
	AirportID uint   `json:"airport_id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Address   string `json:"address"`
}

// ==========================
// Detail untuk GET /airports/:id
// ==========================
type DetailAirportDTO struct {
	AirportID uint                 `json:"airport_id"`
	Name      string               `json:"name"`
	Code      string               `json:"code"`
	Address   string               `json:"address"`
	CreatedAt int64                `json:"created_at"`
	UpdatedAt int64                `json:"updated_at"`
	Users     []SummaryUserDTO     `json:"users"`
	Devices   []SummaryDeviceDTO   `json:"devices"`
	Playlists []SummaryPlaylistDTO `json:"playlists"`
}

// ==========================
// Create (Request & Response)
// ==========================
type CreateAirportReqDTO struct {
	Name    string `json:"name" binding:"required"`
	Code    string `json:"code" binding:"required"`
	Address string `json:"address"`
}

type CreateAirportResDTO struct {
	AirportID uint   `json:"airport_id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ==========================
// Update (Request & Response)
// ==========================
type UpdateAirportReqDTO struct {
	AirportID uint    `json:"airport_id" binding:"required"`
	Name      *string `json:"name,omitempty"`
	Code      *string `json:"code,omitempty"`
	Address   *string `json:"address,omitempty"`
}


type UpdateAirportResDTO struct {
	AirportID uint   `json:"airport_id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
