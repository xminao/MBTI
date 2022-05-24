package main

import "fmt"

func main() {
	/*
	 * Go语言取地址符：&
	 */
	var a int = 10
	fmt.Println(&a)

	/*
	 * 指针声明： var var_name *var-type
	 */
	var ip *int
	ip = &a
	fmt.Println(ip, *ip) //使用*指针变量名访问值

	/*
	 * Go空指针，没有分配任何变量则为nil，nil和其它语言的null none NULL一样，都指为零值
	 */
	var ptr *int
	fmt.Println(ptr)
	fmt.Printf("%x", ptr)

	//空指针判断
	if ptr != nil {
		fmt.Println("非空指针")
	}
	if ptr == nil {
		fmt.Println("空指针")
	}

	/*
	 * 指针指向数组
	 */
	b := []int{1, 2, 4, 5}
	var ptr_2 [4]*int

	for i := 0; i < 4; i++ {
		ptr_2[i] = &b[i] //将a中元素地址赋给指针数组ptr_2
	}

	for i := 0; i < 4; i++ {
		fmt.Println(i, *ptr_2[i])
	}

	//指针作为函数参数
	d := 10
	e := 20
	swap(&d, &e)
	fmt.Println(d, e)
}

/*
 * 指针作为函数参数，也就是之前写的值传递
 */

// 一个更加简单的交换函数
func swap_3(x *int, y *int) {
	*x, *y = *y, *x
}
