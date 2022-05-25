package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"                   //导入包
	"github.com/jinzhu/gorm/dialects/postgres" //导入PgSql驱动程序
)

/* 定义模型 */
type Student_profile struct {
	Sno               string         `gorm:"primary_key;not null"`
	Student_informat  postgres.Jsonb `json:"student_informat"`
	Other_information postgres.Jsonb `json:"other_information"`
}

/*
 * 查询数据
 * json字符串的byte数组反序列化
 * 反序列化：
 */
func toJson(data []byte) map[string]interface{} {
	var ben map[string]interface{}
	json.Unmarshal(data, &ben)
	return ben
}

/* 插入数据 */
func insertData(db *gorm.DB) {
	sno := "{2,0,2,0,1,2}"
	student_informat := postgres.Jsonb{[]byte(`{ "name": "许敏浩", "gender": "男", "age": "21", "politic": "共青团员", "fresh": "2", "education": "本科"}`)}
	other_informat := postgres.Jsonb{[]byte(`{ "source": "大连民族大学", "major": "软件工程", "category": "电子信息"}`)}
	stu := Student_profile{sno, student_informat, other_informat}
	db.Create(&stu)
	fmt.Println("插入成功")
}

/* 查询数据 */
func selectData(db *gorm.DB) {
	var stus []Student_profile
	/* 获取所有记录 */
	db.Find(&stus)

	for _, stu := range stus {
		fmt.Println(stu.Sno, toJson(stu.Student_informat.RawMessage), toJson(stu.Student_informat.RawMessage))
	}
}

func main() {
	/* 连接数据库 */
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gorm_learning sslmode=disable password=admin")

	/*
	 *  在Gorm中，表名是结构体名的复数形式，列名是字段名的蛇形小写。
	 * 如果有一个user表，那么你定义的结构体名为：User，gorm会默认表名为users而不是user
	 * 加上db.SingularTable(true) gorm转义struct名字就不会加上s了！
	 */
	db.SingularTable(true)

	/* 检测连接状态 */
	if err != nil {
		panic(err)
	} else {
		fmt.Println("连接成功")
	}

	if !db.HasTable(&Student_profile{}) {
		db.AutoMigrate(&Student_profile{})
		fmt.Println("建表成功")
	}

	/*  插入数据 */
	insertData(db)

	defer db.Close()
}
