package models

type Log struct {
	LogID     uint   `gorm:"primaryKey;autoIncrement;column:log_id"`
	UserID    *uint  `gorm:"column:user_id"`
	DeviceID  *uint  `gorm:"column:device_id"`
	Timestamp int64  `gorm:"column:timestamp;not null"`
	Source    string `gorm:"size:50;not null;column:source"`
	Action    string `gorm:"size:100;not null;column:action"`
	Message   string `gorm:"type:text;column:message"`

	User   *User   `gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	Device *Device `gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
}
