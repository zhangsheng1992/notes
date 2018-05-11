package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	test1()
	test2()
	test3()
	test4()
}

/**
 *	simple http request
 */
func test1() {
	response, _ := http.Get("http://127.0.0.1:8888/get?a=1&b=2&c=abcd")
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("server端返回的信息为:\n", string(data))
}

/**
 *	demo2
 */
func test2() {
	//create Rquest and set Header
	request, _ := http.NewRequest("GET", "http://127.0.0.1:8888/get", nil)
	request.Header["Accept"] = []string{"text/html,", "q=0.8"}
	request.Header["Accept-Language"] = []string{"zh-CN,zh;q=0.9"}

	//init client and send this request
	client := http.Client{}
	response, _ := client.Do(request)
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("server端返回的信息为:\n", string(data))
}

/**
 * demo3
 */
func test3() {
	request := &http.Request{}
	request.Method = "POST"
	u, _ := url.Parse("http://127.0.0.1:8888/get")
	request.URL = u
	request.Header = map[string][]string{}
	request.Header["Accept"] = []string{"text/html,", "q=0.8"}
	request.Header["Accept-Language"] = []string{"zh-CN,zh;q=0.9"}

	client := http.Client{}
	response, _ := client.Do(request)
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("server端返回的信息为:\n", string(data))
}

/**
 * other attributes sets
 */
func test4() {
	request := &http.Request{
		Header: map[string][]string{},
	}
	request.Method = "POST"
	u, _ := url.Parse("http://127.0.0.1:8888/test")
	request.URL = u
	request.Host = "www.a.com"
	request.RemoteAddr = "www.a.com"

	client := http.Client{}
	response, _ := client.Do(request)
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("server端返回的信息为:\n", string(data))
}
