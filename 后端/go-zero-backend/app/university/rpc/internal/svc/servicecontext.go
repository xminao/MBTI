package svc

import (
	"backend/app/university/rpc/internal/config"
	"backend/app/university/rpc/internal/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config      config.Config
	DbEngin     *gorm.DB
	RedisClient *redis.Redis
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
		return nil
	} else {
		fmt.Printf("连接成功")
	}

	// 自动同步更新表结构
	db.AutoMigrate(&models.CollegeInfo{})

	//err = db.Create(&models.CollegeInfo{
	//	Model:       gorm.Model{},
	//	CollegeId:   1,
	//	CollegeName: "计算机科学与工程",
	//}).Error
	//
	//if err != nil {
	//	panic(err)
	//	return nil
	//} else {
	//	fmt.Printf("插入数据成功")
	//}

	return &ServiceContext{
		Config: c,
		//Gorm
		DbEngin: db,
		//redis
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
