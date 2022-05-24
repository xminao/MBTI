package main

import "fmt"

func main() {
	/*
	 * 算术运算符：
	 * + - * / % ++ --
	 */
	var a int = 21
	var b int = 10
	var c int

	c = a / b //相除
	fmt.Println(c)
	c = a % b //求余
	fmt.Println(c)

	/*
	 * 关系运算符：
	 * == != > < >= <=
	 */
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	/*
	 * 逻辑运算符：
	 * && || !
	 * && 逻辑AND运算符
	 * || 逻辑OR运算符
	 * ！ 逻辑NOT运算符
	 */
	var d = true
	fmt.Println(!d)

	/*
	 * 赋值运算符
	 * = += -+ *= /+ %=
	 */

	/*
	 *  其它运算符
	 * &返回变量存储地址 *指针变量
	 */

}
