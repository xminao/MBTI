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
	}

	var stus []Student_profile
	/* 获取所有记录 */
	db.Find(&stus)

	for _, stu := range stus {
		fmt.Println(stu.Sno, toJson(stu.Student_informat.RawMessage), toJson(stu.Student_informat.RawMessage))
	}

	defer db.Close()
}
