package models

import "github.com/jinzhu/gorm"

type Userinfo struct {
	gorm.Model
	Id       int    `json:"id" gorm:"type:integer;not null;" binding:"required"`
	Name     string `json:"name" gorm:"type:varchar(255);not null;default '';" binding:"required"`
	Password string `json:"password" gorm:"type:varchar(255);not null;default '';" binding:"required"`
	Age      int    `json:"age" gorm:"type:smallint;not null;default 0;" binding:"required"`
	Gender   string `json:"gender" gorm:"type:varchar(5);not null;" binding:"required"`
}
