> body为http请求的一部分,本章将使用golang的设置http请求的body内容

### 一个简单的post请求

一个http请求中，请求正文紧跟在请求行，请求头之后，请求正文中可以包含客户提交数据信息，也可以不包含任何数据

本篇的所有客户端demo与服务端demo可以在xxx下载，来看一个demo:

构造一个request结构并设置url与请求方式，然后在body中传递一个简单的字符串，发送请求

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
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

```

运行结果

服务端返回信息: 服务端收到的数据为: hello world!

### io.ReadCloser

**request.Body** 接收 **io.ReadCloser** 类型的参数，常见的数据一般以[]byte和string为主，

两种数据类型本身是可以互相转换的，所以转换成 **io.ReadCloser** 的方法也差不多

```go
package main

import (
	"bytes"
	"fmt"
	"strings"
)
func main(){
	data := "hello world!"
	reader := strings.NewReader(data)
	readcloser := ioutil.NopCloser(reader)
	fmt.Println(readcloser)
}
```

```go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	data := make([]byte, 10)
	reader := bytes.NewReader(data)
	readcloser := ioutil.NopCloser(reader)
	fmt.Println(readcloser)
}	
```

### 设置form参数

可以通过 request.Body设置post参数

注意:传递的post参数要能被正确解析,需要设置Header

```go
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
```

```go
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
```

运行结果

服务端返回信息: 服务端信息:
key=a,value=1
key=b,value=2
key=c,value=3
key=d,value=4444
key=abc_id,value=test
