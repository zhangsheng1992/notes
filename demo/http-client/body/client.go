package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//	test1()
	test2()
}

/**
 *	设置body为一个字符串
 */
func test1() {
	//构造request结构
	request := &http.Request{}
	url, _ := url.Parse("http://127.0.0.1:8888/test")
	request.URL = url
	request.Method = "POST"

	//将字符串转换为io.ReadCloser类型
	data := "hello world!"
	reader := strings.NewReader(data)
	request.Body = ioutil.NopCloser(reader)

	//发送请求并解析返回的内容
	response, _ := new(http.Client).Do(request)
	result, _ := ioutil.ReadAll(response.Body)

	fmt.Println("服务端返回信息:", string(result))
}

/**
 *	利用request.Body 直接设置请求的post参数
 */
func test2() {
	request := &http.Request{}
	url, _ := url.Parse("http://127.0.0.1:8888/test2")
	request.URL = url
	request.Method = "POST"

	//设置请求header 使server可以解析post参数
	request.Header = http.Header{}
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	data := "a=1&b=2&c=3&d=4444&abc_id=test"
	reader := strings.NewReader(data)
	request.Body = ioutil.NopCloser(reader)

	//发送请求并解析返回的内容
	response, _ := new(http.Client).Do(request)
	result, _ := ioutil.ReadAll(response.Body)

	fmt.Println("服务端返回信息:", string(result))
}
