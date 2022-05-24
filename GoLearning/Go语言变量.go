package main

import "fmt"

/*
 * 数据类型：包括bool、int、float32、float64、string等
 */

func main() {
	/*
	 * Go变量名由 字母、数字、下划线构成
	 * 变量声明使用var关键字，可一次声明多个变量
	 * 格式： var identifier type / var identifier1, identifier2 type
	 *
	 */
	var a string = "xminao"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	/*
	 * 指定变量名，但没有初始化则变量默认为零值
	 */
	var d int
	fmt.Println(d)

	// 默认输出false
	var f bool
	fmt.Println(f)

	/*
	 * 也可以根据值自己判断
	 */
	var g = true
	fmt.Println(g)

	/*
	 * := 声明语句进行声明和初始化
	 */
	intVal := 1
	// 相当于
	var intVal_1 int
	intVal_1 = 1

	fmt.Println(intVal)
	fmt.Println(intVal_1)
}
