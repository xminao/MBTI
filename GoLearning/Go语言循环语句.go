package main

import "fmt"

//搞清楚fmt.Printf 和 fmt.Println的区别！！！

/*
 * Go语言提供两种类型循环处理语句：
 * for循环、循环嵌套
 * 循环控制语句：
 * break、continue、goto
 */

func main() {
	/*
	 * Go语言for循环有三种形式，其中一种使用分号
	 * for init; condition; post {} 和C语言for一样
	 * for condition {} 和C语言while一样
	 * for {} 和C语言for(;)一样 range
	 */

	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1-10和为", sum)

	/*
	 * for的range格式可以对slice、map、数组、字符串进行迭代循环
	 * for key, value := range oldMap {
	 * 	newMap[key] = value
	 * }
	 * 如果只读key：
	 * for key := range oldMap / for key, _ := range oldMap
	 * 如果只读value：
	 * for _, value := range oldMap
	 */

	map1 := make(map[int]float32) //使用内置的make函数创建map
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取key value
	for key, value := range map1 {
		fmt.Printf("key:%d, value:%f", key, value)
	}

	//读取key
	for key, _ := range map1 {
		fmt.Printf("key:%d", key)
	}

	//读取value
	for _, value := range map1 {
		fmt.Printf("value: %f", value)
	}

	//读取key和value

	// for-each range 循环
	strings := []string{"许敏浩", "许敏浩2", "许敏浩3"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
}
