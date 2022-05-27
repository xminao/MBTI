package models

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type QuestionBank struct {
	gorm.Model
	//Id           int64          `gorm:"type:integer;not null;"`
	QuestionId   int64          `gorm:"type:integer;primary_key;not null;"`
	QuestionType string         `gorm:"type:varchar(255);not null"`
	QuestionText postgres.Jsonb `json:"question_text"`
}
