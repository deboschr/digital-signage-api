package models

type Airport struct {
    AirportID uint   `gorm:"primaryKey;column:airport_id"`
    Name      string `gorm:"size:150;not null;column:name"`
    Code      string `gorm:"size:10;unique;not null;column:code"`
    Address   string `gorm:"size:255;column:address"`
    CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at"`
    UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at"`
}