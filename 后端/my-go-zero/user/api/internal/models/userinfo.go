package models

import "github.com/jinzhu/gorm"

type Userinfo struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null;default '';"`
	Password string `gorm:"type:varchar(255);not null;default '';"`
	Age      int64  `gorm:"type:smallint;not null;default 0;"`
	Gender   string `gorm:"type:varchar(5);not null;"`
	// gorm不需要自己建这些列
	//Id       int64  `gorm:"type:integer;not null;"`
	//CreateTime int64  `gorm:"autoCreateTime"`
	//UpdateTime int64  `gorm:"autoUpdateTime"`
}
