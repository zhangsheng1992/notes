package main

import (
	"fmt"
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
	http.HandleFunc("/test", test)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("start server error:", err)
	}
}

/**
 *	a simple http handle function
 *	it will print the request's head
 */
func simpleGet(w http.ResponseWriter, r *http.Request) {
	str := ""
	for key, value := range r.Header {
		str = str + key + "=" + strings.Join(value, "") + "\n"
	}
	fmt.Fprintln(w, str)
}

/**
 *	a simple http handle function
 *	it will print the request's head
 */
func test(w http.ResponseWriter, r *http.Request) {
	str := "Host:" + r.Host + "\n" + "RemoteAddr:" + r.RemoteAddr + "\n"
	fmt.Fprint(w, str)
}
