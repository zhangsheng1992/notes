package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	//get()
	post()
}

/**
 *	a simple http post request
 */
func post() {
	response, err := http.Post("http://127.0.0.1:8888/post",
		"application/x-www-form-urlencoded", strings.NewReader("a=1&b=2"))
	if err != nil {
		fmt.Println("reqeust error:", err)
	}
	data, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println("read response Body error:", readErr)
	} else {
		fmt.Println(string(data))
	}
}

/**
 *	a simple http get request
 */
func get() {
	response, err := http.Get("http://127.0.0.1:8888/get?a=1&b=2&c=abcd")
	if err != nil {
		fmt.Println("reqeust error:", err)
	}
	data, readErr := ioutil.ReadAll(response.Body)
	if readErr != nil {
		fmt.Println("read response Body error:", readErr)
	} else {
		fmt.Println(string(data))
	}
}
