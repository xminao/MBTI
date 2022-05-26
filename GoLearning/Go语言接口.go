package main

import "fmt"

/*
 * 把所有具有共性的方法定义在一起
 * 其他类型只要实现了这些方法就是实现了这个接口
 */

/*
	// 定义接口
	type interface_name inteface {
		method_name [return_type]
		method_name2 [return_type]
	}

	// 定义结构体
	type struct_name struct {
		//方法实现
	}

	// 实现接口方法
	func (struct_name_variable struct_name) method_name() [return_type] {
			//方法实现
	}
*/

// 定义了接口Phone，接口里有一个方法叫call
type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("我是诺基亚，call")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("我是苹果，call")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
