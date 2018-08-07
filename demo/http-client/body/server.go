package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 创建一个http服务器 监听8888端口.如果启动出错,检查一下端口是否被占用。
func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/test2", test2)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("start server error:", err)
	}
}

// 这个路由处理器直接打印出请求的body
func test(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintln(w, "服务端收到的数据为:", string(data))
}

// 打印post请求参数
func test2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	str := "服务端信息:\n"
	for k, v := range r.PostForm {
		str = str + "key=" + k + ",value=" + v[0] + "\n"
	}
	fmt.Fprintln(w, str)
}
