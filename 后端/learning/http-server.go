/*
 * 使用http包的Server启动一个http服务，两个方式
 */
package main

import "net/http"

/* 第一种实现方法 */
//func main() {
//	/*
//	 * *http.Request 是http请求内容实例的指针
//	 * http.ResponseWriter 些http相应内容的实例
//	 */
//	http.HandleFunc("/v1/demo", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("Hello World\n"))
//	})
//
//	// 启动一个htttp服务并监听8888端口，第二个参数可以指定handler
//	http.ListenAndServe(":8888", nil)
//
//}

/*
 * 上面的ListenAndServe 是对http.Server的进一步封装
 * 可以使用http.Server直接启动服务，需要设置Handler
 * 这个Handler要实现Server.Handler这个接口
 * 当收到请求会执行这个Handler的ServeHTTP方法
 */

/* 第二种实现方法 */

//DemoHandle server handle实例
type DemoHandle struct {
	Name string
	Sex  string
}

// ServeHTTP 匹配到路由后执行的方法
func (demoHandle *DemoHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World\n"))
}

func main() {
	http.Handle("/", &DemoHandle{Name: "许敏浩", Sex: "男"})
	http.ListenAndServe(":8889", nil)
}
