package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/**
 *	create a http server to provide services
 *	if the error happened,usually,the reason of
 *	the error is the 8080 port be occupyed by other application.
 *	please check and try again.
 */
func main() {
	http.HandleFunc("/get", simpleGet)
	http.HandleFunc("/post", simplePost)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("start server error:", err)
	}
}

/**
 *	a simple http handle function
 *	it will print the request's parameters
 *	visit http://127.0.0.1:8888/get?a=1&b=2&b=3
 *	for more infomation
 * 	the same of parameter's name will be wrote to the slice
 */
func simpleGet(w http.ResponseWriter, r *http.Request) {
	var str string
	for k, v := range r.URL.Query() {
		if len(v) == 1 {
			str = str + fmt.Sprintln(k, "=", v[0])
		} else {
			//join by ","
			str = str + fmt.Sprintln(k, "=", strings.Join(v, ","))
		}
	}
	fmt.Fprint(w, str)
}

/**
 *	a simple http handle function like the above handle function
 *	it will print the request's body if the request method is post
 */
func simplePost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	} else {
		fmt.Fprintln(w, "the request body is "+string(data))
	}
}
