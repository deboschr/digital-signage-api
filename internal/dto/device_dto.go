package dto

type GetSummaryDeviceResDTO struct {
	DeviceID    uint   `json:"device_id"`
	Name        string `json:"name"`
	IsConnected string `json:"is_connected"`
}

type GetDetailDeviceResDTO struct {
	DeviceID    uint                    `json:"device_id"`
	Name        string                  `json:"name"`
	ApiKey      string                  `json:"api_key"`
	IsConnected string                  `json:"is_connected"`
	Airport     GetSummaryAirportResDTO `json:"airport"`
}

type CreateDeviceReqDTO struct {
	AirportID uint   `json:"airport_id" binding:"required,gt=0"`
	Name      string `json:"name" binding:"required,min=3,max=100"`
	ApiKey    string `json:"api_key" binding:"required,len=64"`
}

type UpdateDeviceReqDTO struct {
	DeviceID  uint    `json:"device_id" binding:"required,gt=0"`
	AirportID *uint   `json:"airport_id" binding:"omitempty,gt=0"`
	Name      *string `json:"name" binding:"omitempty,min=3,max=100"`
	ApiKey    *string `json:"api_key" binding:"omitempty,len=64"`
}
