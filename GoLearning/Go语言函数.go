package main

import (
	"fmt"
	"math"
)

/*
 * Go语言至少有一个main()函数
 * 格式也挺不一样：
 * func function_name([parameter list]) [return_types] { 函数体 }
 */

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, b)
	fmt.Printf("最大值：%d", ret)

	/*
	 * 函数还可以返回多个值
	 */
	c, d := swap("前", "后")
	fmt.Println(c, d)

	/*
	 * 引用传递
	 */
	swap_2(&a, &b)
	fmt.Printf("a:%d, b:%d\n", a, b)

	/*
	 * Go语言中，函数可以作为函数的实参。
	 * 比如定义一个函数 getSquareRoot并在Println函数中输出
	 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}

func swap(x, y string) (string, string) {
	return y, x
}

/*
 * 可以使用 值传递、引用传递来传递函数
 * 值传递：
 * 		将实际参数复制一份传递到函数，Go语言默认使用的是值传递，在调用过程中不会影响实际参数
 * 引用传递：
 *		将实际参数的地址传递到函数中，如果在函数中对参数进行修改，则将影响到实际参数
 */

func swap_2(x *int, y *int) { //定义指针变量x 和指针变量y
	var temp int
	// 和C语言指针一样
	temp = *x //这里*x是取该地址的值，也就是说取得了传进来参数的值所在地址然后对值进行操作
	*x = *y
	*y = temp
}
