package models

type User struct {
	UserID    uint   `gorm:"primaryKey;autoIncrement;column:user_id"`
	AirportID *uint  `gorm:"column:airport_id"` // nullable foreign key
	Username  string `gorm:"size:100;unique;not null;column:username"`
	Password  string `gorm:"size:255;not null;column:password"`
	Role      string `gorm:"type:enum('management','admin','operator');not null;column:role"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at"`

	Airport Airport `gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
}
