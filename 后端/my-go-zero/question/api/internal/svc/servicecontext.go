package svc

import (
	"github.com/jinzhu/gorm"
	"question/api/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//// 连接数据库
	//dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	c.Database.Host,
	//	c.Database.Port,
	//	c.Database.User,
	//	c.Database.Password,
	//	c.Database.Name,
	//)
	//
	//db, err := gorm.Open("postgres", dsn)
	//
	//db.SingularTable(true)
	//
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Printf("连接成功")
	//}
	//
	//// 自动同步更新表结构
	//db.AutoMigrate(&models.QuestionBank{})

	//test := models.QuestionBank{
	//	Model:        gorm.Model{},
	//	QuestionId:   1,
	//	QuestionType: "A",
	//	QuestionText: postgres.Jsonb{[]byte(`{"question_desc": "测试", "option_a": "选项A", "option_b": "选项B"}`)},
	//}
	//
	//err_insert := db.Create(&test).Error
	//if err_insert != nil {
	//	panic(err_insert)
	//} else {
	//	fmt.Printf("数据插入成功")
	//}

	//return &ServiceContext{
	//	Config:  c,
	//	DbEngin: db,
	//}
	l.
}
