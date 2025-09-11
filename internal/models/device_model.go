package models

type Device struct {
	DeviceID  uint   `gorm:"primaryKey;autoIncrement;column:device_id"`
	AirportID uint   `gorm:"not null;column:airport_id"`
	Name      string `gorm:"size:100;not null;column:name"`
	IpAddress string `gorm:"size:50;unique;column:ip_address"`
	Status    string `gorm:"size:20;column:status"`
	CreatedAt int64  `gorm:"autoCreateTime:milli;column:created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli;column:updated_at"`

	Airport *Airport `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
}
