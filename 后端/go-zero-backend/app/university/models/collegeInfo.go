package models

type CollegeInfo struct {
	CollegeId   int    `gorm:"type:serial;not null;primary_key;"`
	CollegeName string `gorm:"type:varchar(255);not null;"`
}
