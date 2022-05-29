package svc

import (
	"backend/app/university/api/internal/config"
	"backend/app/university/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
	)
	db, err := gorm.Open("postgres", dsn)

	db.SingularTable(true)

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("连接成功")
	}

	// 自动同步更新表结构
	db.AutoMigrate(&models.CollegeInfo{}, &models.YearInfo{}, &models.MajorInfo{}, &models.ClassInfo{}, &models.StudentInfo{})

	return &ServiceContext{
		Config: c,
		//Gorm
		DbEngin: db,
	}
}
