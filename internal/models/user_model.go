package models

type User struct {
	UserID       uint   `gorm:"primaryKey;column:user_id"`
	Username     string `gorm:"size:100;unique;not null;column:username"`
	PasswordHash string `gorm:"size:255;not null;column:password_hash"`
	Role         string `gorm:"type:enum('management','admin','operator');not null;column:role"`
	CreatedAt    int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli;column:updated_at"`
}