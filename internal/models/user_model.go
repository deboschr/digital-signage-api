package models

type User struct {
	UserID    uint   `gorm:"primaryKey;autoIncrement;column:user_id"`
	AirportID *uint  `gorm:"column:airport_id"`
	Username  string `gorm:"size:100;unique;not null;column:username"`
	Password  string `gorm:"size:255;not null;column:password"`
	Role      string `gorm:"type:enum('admin','operator','management');not null;column:role"`

	Airport *Airport `gorm:"constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
	Logs    []*Log   `gorm:"foreignKey:UserID;constraint:OnDelete:SET NULL,OnUpdate:CASCADE"`
}
