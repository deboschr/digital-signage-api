package models

type User struct {
    UserID    uint   `gorm:"primaryKey;column:user_id"`
    AirportID *uint  `gorm:"column:airport_id"` // nullable, supaya management bisa null
    Username  string `gorm:"size:100;unique;not null;column:username"`
    Password  string `gorm:"size:255;not null;column:password"`
    Role      string `gorm:"type:enum('management','admin','operator');not null;column:role"`
    CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at"`
    UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at"`

    Airport Airport `gorm:"foreignKey:AirportID"`
}
