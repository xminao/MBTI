package main

import "fmt"

/*
 * 切片：Slice
 * Go数组长度不可改变，Go提供了切片，与数组相比，切片长度不固定，可以追加元素，追加元素可以使切片容量增大
 * 声明一个未指定大小的数组来定义切片：
 * var identifier []type 切片不需要说明长度
 * 使用make函数创建切片，len是数组的长度也是切片的初始长度
 * var slice1 []type = make([]type, len)
 * 或简写： slice1 := make([]type, len)
 * 可以指定容量，加一个capacity参数（可选）
 * make([]T, length, capacity)
 *
 * 初始化切片：
 * s := []int{1, 2, 3}
 * s := arr[:] / s := arr[startIndex:endIndex] 切片s是数组arr的引用
 *
 * len()方法获取切片长度，cap()测量切片最长可以达到多少
 */

func main() {
	// 空切片，未初始化之前默认为nil
	var numbers = make([]int, 3, 5)
	printSlice(numbers)

	//切片截取 ,左闭右开[:8]取index为0-7的，[3:]取index为3-最后的元素
	number_2 := []int{1, 3, 4, 6, 7, 8, 43, 56}
	printSlice(number_2)
	fmt.Println(number_2[3:6])
	fmt.Println(number_2[:8])
	fmt.Println(number_2[3:])

	/*
	 * append()追加新元素
	 * copy()拷贝切片
	 */
	//空切片
	var number_3 []int
	printSlice(number_3)
	number_3 = append(number_3, 2)
	printSlice(number_3)
	number_3 = append(number_3, 4, 5, 6) //追加多个元素
	printSlice(number_3)

	//创建一个两倍number_3容量的切片
	number_4 := make([]int, len(number_3), (cap(number_3) * 2))
	copy(number_4, number_3) //拷贝后面的到前面
	printSlice(number_4)
}

func printSlice(x []int) {
	fmt.Println(len(x), cap(x), x)
}
