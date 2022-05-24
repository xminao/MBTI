package main

import "fmt"

const (
	Unkown = 0
	Female = 1
	Male   = 2
)

func main() {
	/*
	 * 常量顾名思义，在程序运行时候不能改变的值
	 * 定义：const identifier [type] = value
	 */
	const a string = "abc"    //显式类型定义
	const b = "abc"           //隐式类型定义
	const c, d = "bcd", "efg" //多个同类型定义声明简写

	/*
	 * 常量可以枚举
	 * 常量可以用内置函数，否则编译不通过
	 */
	fmt.Println(Unkown, Female, Male)

}
