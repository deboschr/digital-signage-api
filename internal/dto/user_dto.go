package dto

type GetSummaryUserResDTO struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type GetDetailUserResDTO struct {
	UserID   uint                     `json:"user_id"`
	Username string                   `json:"username"`
	Role     string                   `json:"role"`
	Airport  *GetSummaryAirportResDTO `json:"airport,omitempty"`
}

type CreateUserReqDTO struct {
	AirportID *uint  `json:"airport_id" binding:"omitempty,gt=0"`
	Username  string `json:"username" binding:"required,min=3,max=100"`
	Password  string `json:"password" binding:"required,min=6,max=255"`
	Role      string `json:"role" binding:"required,oneof=management admin"`
}

type UpdateUserReqDTO struct {
	UserID    uint    `json:"user_id" binding:"required,gt=0"`
	AirportID *uint   `json:"airport_id" binding:"omitempty,gt=0"`
	Username  *string `json:"username" binding:"omitempty,min=3,max=100"`
	Password  *string `json:"password" binding:"omitempty,min=6,max=255"`
	Role      *string `json:"role" binding:"omitempty,oneof=management admin"`
}
