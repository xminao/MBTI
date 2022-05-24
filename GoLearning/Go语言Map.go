package main

import "fmt"

/*
 * Map是一种无需的键值对的集合，通过key快速检索数据，key类似索引，指向数据的值
 * Mao是一种集合，所以可以像迭代数组一样迭代，不过它是无序的，无法决定返回顺序，因为map是使用hash实现的
 */

func main() {
	/*
	 * 定义Map
	 * var map_variable map[key_data_type]value_data_type
	 * 或者使用make函数
	 * map_variable := make(map[key_data_type]value_data_type)
	 * 如果不初始化，就会创建一个nil map，无法存放键值对
	 */

	//var capitalMap map[string]string
	//capitalMap = make(map[string]string) //初始化
	capitalMap := make(map[string]string) //直接声明加初始化

	capitalMap["France"] = "巴黎"
	capitalMap["Italy"] = "罗马"
	capitalMap["Japan"] = "东京"

	for country := range capitalMap {
		fmt.Println(country, "首都是：", capitalMap[country])
	}

	capital, ok := capitalMap["China"]
	if ok {
		fmt.Println("存在", capital)
	} else {
		fmt.Println("不存在")
	}

	//delete()函数用于删除集合的元素，参数为map和对应的key
	delete(capitalMap, "France")
	fmt.Println(capitalMap)
}
