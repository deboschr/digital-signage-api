package dto

// ==========================
// Summary untuk listing
// ==========================
type SummaryDeviceDTO struct {
	DeviceID  uint   `json:"device_id"`
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
	Status    string `json:"status"`
}

// ==========================
// Detail untuk GET /devices/:id
// ==========================
type DetailDeviceDTO struct {
	DeviceID  uint               `json:"device_id"`
	Name      string             `json:"name"`
	IpAddress string             `json:"ip_address"`
	Status    string             `json:"status"`
	CreatedAt int64              `json:"created_at"`
	UpdatedAt int64              `json:"updated_at"`
	Airport   *SummaryAirportDTO `json:"airport,omitempty"`
}

// ==========================
// Create (Request & Response)
// ==========================
type CreateDeviceReqDTO struct {
	AirportID uint   `json:"airport_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	IpAddress string `json:"ip_address" binding:"required"`
	Status    string `json:"status"`
}

type CreateDeviceResDTO struct {
	DeviceID  uint   `json:"device_id"`
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ==========================
// Update (Request & Response)
// ==========================
type UpdateDeviceReqDTO struct {
	DeviceID  uint    `json:"device_id" binding:"required"`
	AirportID *uint   `json:"airport_id,omitempty"`
	Name      *string `json:"name,omitempty"`
	IpAddress *string `json:"ip_address,omitempty"`
	Status    *string `json:"status,omitempty"`
}


type UpdateDeviceResDTO struct {
	DeviceID  uint   `json:"device_id"`
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
	Status    string `json:"status"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
