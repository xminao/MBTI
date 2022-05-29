package models

type ClassInfo struct {
	ClassId     int         `gorm:"type:serial;not null;primary_key;"`
	ClassName   string      `gorm:"type:varchar(255);not null;"`
	CollegeInfo CollegeInfo `gorm:"ForeignKey:CollegeID;AssociationForeignKey:CollegeID"`
	CollegeId   int
	YearInfo    YearInfo `gorm:"ForeignKey:Year;AssociationForeignKey:Year"`
	Year        int
	MajorInfo   MajorInfo `gorm:"ForeignKey:MajorId;AssociationForeignKey:MajorId"`
	MajorId     int
}
