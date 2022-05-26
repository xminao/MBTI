// HTTP SERVER
package main

import (
	"net/http"
)

func main() {
	/*
	 * *http.Request 是http请求内容实例的指针
	 * http.ResponseWriter 些http相应内容的实例
	 */
	http.HandleFunc("/v1/demo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World\n"))
	})

	// 启动一个htttp服务并监听8888端口，第二个参数可以指定handler
	http.ListenAndServe(":8888", nil)
}
