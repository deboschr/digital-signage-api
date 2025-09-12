package dto

// ==========================
// Summary untuk listing
// ==========================
type SummaryUserDTO struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// ==========================
// Detail untuk GET /users/:id
// ==========================
type DetailUserDTO struct {
	UserID    uint              `json:"user_id"`
	Username  string            `json:"username"`
	Role      string            `json:"role"`
	CreatedAt int64             `json:"created_at"`
	UpdatedAt int64             `json:"updated_at"`
	Airport   *SummaryAirportDTO `json:"airport,omitempty"`
}

// ==========================
// Create (Request & Response)
// ==========================
type CreateUserReqDTO struct {
	AirportID *uint  `json:"airport_id,omitempty"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
}

type CreateUserResDTO struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// ==========================
// Update (Request & Response)
// ==========================
type UpdateUserReqDTO struct {
	UserID    uint    `json:"user_id" binding:"required"`
	AirportID *uint   `json:"airport_id,omitempty"`
	Username  *string `json:"username,omitempty"`
	Password  *string `json:"password,omitempty"`
	Role      *string `json:"role,omitempty"`
}


type UpdateUserResDTO struct {
	UserID    uint   `json:"user_id"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
