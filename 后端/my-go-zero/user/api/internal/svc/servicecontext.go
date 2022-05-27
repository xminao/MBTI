package svc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"user/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//UserModel model.UserinfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 连接数据库
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
	)
	// 注意是postgres
	db, err := gorm.Open("postgres", dsn)
	db.SingularTable(true)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("连接成功")
	}

	//db.AutoMigrate(&models.User{})

	return &ServiceContext{
		Config: c,
		//UserModel: model.NewUserinfoModel(db),
	}
}
