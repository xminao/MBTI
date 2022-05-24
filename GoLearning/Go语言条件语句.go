package main

import "fmt"

/*
 * Go语言提供了几种条件判断语句：
 * if
 * if... else
 * if 嵌套语句
 * switch 语句
 * select 语句
 */

func main() {
	/*
	 * switch比较特别
	 * 从上至下执行，直到找到匹配项，匹配项后不需要加break
	 * switch 默认情况下case后自带break，匹配成功不会执行别的case
	 * 当我们你需要执行后面的case，可以加fallthrough关键字
	 * 添加fallthrough关键字后，直到没有fallthrough的case语句才会停止
	 */

	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
		fallthrough
	case false:
		fmt.Println("false")
		fallthrough
	case true:
		fmt.Println("true")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("default")
	}
}
