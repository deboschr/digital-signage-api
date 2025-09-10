package models

import "time"

type Device struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:100;not null"`
    IPAddress string    `gorm:"size:50;unique"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
