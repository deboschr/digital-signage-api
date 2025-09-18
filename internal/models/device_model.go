package models

type Device struct {
	DeviceID    uint   `gorm:"primaryKey;autoIncrement;column:device_id"`
	AirportID   uint   `gorm:"not null;column:airport_id"`
	Name        string `gorm:"size:100;not null;column:name"`
	ApiKey      string `gorm:"size:64;not null;unique;column:api_key"` // secret unik per device
	IsConnected bool   `gorm:"type:tinyint(1);not null;default:0;column:is_connected"`

	Airport *Airport `gorm:"constraint:OnDelete:RESTRICT,OnUpdate:CASCADE"`
	Logs    []*Log   `gorm:"foreignKey:DeviceID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
}
