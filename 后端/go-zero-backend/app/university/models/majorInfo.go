package models

type MajorInfo struct {
	MajorId     int         `gorm:"type:serial;not null;primary_key;"`
	MajorName   string      `gorm:"type:varchar(255);not null;"`
	CollegeInfo CollegeInfo `gorm:"ForeignKey:CollegeID;AssociationForeignKey:CollegeID"`
	CollegeId   int
	YearInfo    YearInfo `gorm:"ForeignKey:Year;AssociationForeignKey:Year"`
	Year        int
}
