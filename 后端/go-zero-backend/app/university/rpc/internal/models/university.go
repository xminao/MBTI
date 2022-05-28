package models

import "github.com/jinzhu/gorm"

// CollegeInfo 如果不用帕斯卡命名，生成的就是非蛇形，这个将生成college_info的表名
type CollegeInfo struct {
	//gorm的model
	gorm.Model
	CollegeId   int    `gorm:"type:integer;not null;primary_key;"`
	CollegeName string `gorm:"type:varchar(255);not null;"`
}
