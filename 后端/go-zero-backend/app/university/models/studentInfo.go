package models

type StudentInfo struct {
	StudentId   string      `gorm:"type:varchar(255);not null;primary_key;"`
	StudentName string      `gorm:"type:varchar(255);not null;"`
	CollegeInfo CollegeInfo `gorm:"ForeignKey:CollegeID;AssociationForeignKey:CollegeID"`
	CollegeId   int
	YearInfo    YearInfo `gorm:"ForeignKey:Year;AssociationForeignKey:Year"`
	Year        int
	MajorInfo   MajorInfo `gorm:"ForeignKey:MajorId;AssociationForeignKey:MajorId"`
	MajorId     int
	ClassInfo   ClassInfo `gorm:"ForeignKey:ClassId;AssociationForeignKey:ClassId"`
	ClassId     int
}
