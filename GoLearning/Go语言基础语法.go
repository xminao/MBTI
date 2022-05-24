package main

import "fmt"

func main() {
	/*
	 * 标记：关键字、标识符、常量、字符串、符号
	 */
	fmt.Println("标记")

	/*
	 * 字符串连接用 +
	 */
	fmt.Println("连接" + "测试")

	/*
	 * 格式化字符串用 fmt.Sprintf
	 * fmt.Sprintf(格式化样式，参数列表...)
	 */
	var code = 123
	var date = "2022-5-24"
	var url = "Code=%d Date=%s"
	var target_url = fmt.Sprintf(url, code, date)
	fmt.Println(target_url)

}
