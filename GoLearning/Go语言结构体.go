package main

import "fmt"

/*
 * 结构体定义
 * type struct_variable_type struct {
		member definition
		member definition
 * }
 * 一旦定义结构体，就可以用于变量的声明，格式如下：
 * variable_name := struct_variable_type {value1, value2 .. valuen}
 * 或者：
 * variable_name := struct_variable_type {key1:value1, key2:value2 .. keyn:valuen}
*/

type Book struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	// 定义一个结构体,如果忽略不写某字段，则为0
	books := Book{"许敏浩传", "许敏浩", "人文", 1}
	fmt.Println(books)

	/*
	 * 访问结构体成员：
	 * 结构体.成员名
	 */
	books.title = "许敏浩传2"
	books.book_id = 2
	fmt.Println(books)

	//结构体作为函数参数
	printBook(books)

	/*
	 * 结构体指针：
	 * var struct_pointer *Book
	 * 给其赋值 struct_pointer = &books
	 * 使用结构体指针访问结构体成员用"."
	 */
}

/*
 * 结构体作为函数参数
 */
func printBook(book Book) {
	fmt.Println(book)
}
