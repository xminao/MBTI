package models

type YearInfo struct {
	Year        int         `gorm:"type:integer;not null;primary_key"`
	CollegeInfo CollegeInfo `gorm:"ForeignKey:CollegeID;AssociationForeignKey:CollegeID"`
	CollegeId   int
}
