package main

import "fmt"

/*
 * 局部变量：函数内定义的变量
 * 全局变量：函数外定义的变量
 * 形式参数：函数定义中的变量
 */

//声明全局变量
var g int

func main() {
	//声明局部变量
	var a, b, c int

	a = 1
	b = 2
	c = 3
	g = a + b + c

	fmt.Println(a, b, c, g)
}

// a和b就是形参
func sum(a, b int) int {
	return a + b
}
