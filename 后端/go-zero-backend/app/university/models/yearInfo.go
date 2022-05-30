package models

import "github.com/jinzhu/gorm"

type YearInfo struct {
	gorm.Model
	Year        int         `gorm:"type:integer;not null;primary_key"`
	CollegeInfo CollegeInfo `gorm:"ForeignKey:CollegeID;AssociationForeignKey:CollegeID"`
	CollegeId   int
}
