package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 *	create a http server to provide services
 *	if the error happened,usually,the reason of
 *	the error is the 8080 port be occupyed by other application.
 *	please check and try again.
 */
func main() {
	http.HandleFunc("/test", test)
	http.HandleFunc("/test2", test2)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("start server error:", err)
	}
}

/**
 *	a simple http handle function
 *	it will print the request's body
 */
func test(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintln(w, "服务端收到的数据为:", string(data))
}

/**
 *	print post parameter
 */
func test2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	str := "服务端信息:\n"
	for k, v := range r.PostForm {
		str = str + "key=" + k + ",value=" + v[0] + "\n"
	}
	fmt.Fprintln(w, str)
}
