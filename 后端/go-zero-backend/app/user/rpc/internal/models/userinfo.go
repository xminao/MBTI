package models

import "github.com/jinzhu/gorm"

type Userinfo struct {
	gorm.Model
	Username string `gorm:"type:varchar(255);not null;default '';"`
	Nickname string `gorm:"type:varchar(255);not null;default '';"`
	Password string `gorm:"type:varchar(255);not null;default '';"`
	Gender   string `gorm:"type:varchar(5);not null;default '';"`
}
