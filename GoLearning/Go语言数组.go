package main

import "fmt"

/*
 * Go语言数组声明：
 * var variable_name [SIZE] variable_type
 */

func main() {
	//声明一个数组，默认都为0
	var balance [10]float32
	//初始化数组的几种方式
	var balance_1 = [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	balance_2 := [5]float32{1.0, 2.0, 3.0, 4.0, 5.0}
	//如果数组长度不确定，可以使用...代替数组的长度，编译器会自动根据元素个数判断数组的长度
	var balance_3 = [...]float32{3.0, 4.0, 6.0}

	fmt.Println(balance)
	fmt.Println(balance_1)
	fmt.Println(balance_2)
	fmt.Println(balance_3)

	//还可以指定下表初始化元素，要记住是0索引开始
	balance_4 := [5]float32{0: 1.0, 3: 4.0}
	fmt.Println(balance_4)

	/*
	 * 多维数组，和C语言差不多
	 * 创建数组后可以通过append函数向空二维数组添加两行一维数组
	 *
	 */
	values := [][]int{}
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	fmt.Println(values[0])    //输出第一行
	fmt.Println(values[1])    //输出第二行
	fmt.Println(values[0][0]) //输出第一个元素

	//初始化二维数组用大括号来初始值
	a := [3][4]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11}, //这里必须有逗号 或者把下一行的花括号放到这一行就可以不用写逗号
	}

	fmt.Println(a)

	//创建各个维度元素数量不一致的多维数组
	animals := [][]string{}
	row_1 := []string{"fish", "shark", "eel"}
	row_2 := []string{"bird"}
	row_3 := []string{"lizard", "salamander"}

	animals = append(animals, row_1)
	animals = append(animals, row_2)
	animals = append(animals, row_3)

	fmt.Println(animals)
}

/*
 * 向函数传递数组参数，需要在函数定义时声明形参为数组，有两种方式
 */

func myfunc(param [10]int) {

}

// 未确定长度
func myfunc_2(param []int) {

}
