package main

import (
	"fmt"
)

/*
 * 类型转换格式：
 * type_name(expression)
 */

func main() {
	sum := 17
	count := 2
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Println(mean)
}
