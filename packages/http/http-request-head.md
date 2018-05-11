> head为http请求的一部分,下面将介绍如何使用golang内置的http库构造请求头

### http简介
一个完成的http请求由请求行,请求头与请求正文组成

**请求行** 一般用来设置http请求的方式,url与协议版本,一个简单的demo如下:

**`GET/index.html HTTP/1.1`**

表名这是一个GET请求 请求的url为 /index.html 请求的协议为 http1.1

**请求头** 一般用来描述一些客户端相关的信息,如:客户端浏览器版本型号,期望返回的数据类型,期望返回的字符集等等,下面是一个请求的head部分

```go
Accept: text/css,*/*;q=0.1
Accept-Encoding: gzip, deflate, br
Accept-Language: zh-CN,zh;q=0.9
Cache-Control: no-cache
Connection: keep-alive
Host: bkssl.bdimg.com
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.139 Safari/537.36
```

Accept表示客户端期待server返回类型的数据
Accept-Encoding表示客户端希望返回数据的encode方式为gzip(一种压缩格式)
Connection表示客户端希望与服务器进行一段时间的长链接
User-Agent表示客户端的浏览器内核,版本,操作平台等相关信息

**请求正文** 一般包含一些请求的参数,上传的文件信息等等,后续会单独介绍请求正文相关

### 一个简单的demo

server端的程序代码可以点击此处

现在来利用http.Get()方法构造一个简单的http请求

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("http://127.0.0.1:8888/get?a=1&b=2&c=abcd")
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println("server端返回的信息为:\n", string(data))
}
```
运行,将返回:
server端返回的信息为: 
User-Agent=Go-http-client/1.1
Accept-Encoding=gzip

可以看到,通过http.Get()方法请求,server端收到的head信息只有两条,第一行表明了我们是通过 **Go-http-client/1.1** 这个go工具库来请求的,第二行表明了期待的数据encode格式为gzip.

实际使用中，server端可能需要更多的信息，如说明期待请求什么类型的数据，此外一些server端会有部分限制，只处理正常浏览器发出的请求而禁止诸如 **Go-http-client/1.1**这类http client工具的请求，所以有时候就需要自己来构造请求的head了


### 构造请求head
之前的源码中已经介绍了Requset结构中包含了head相关的信息，构造请求head，就需要了解Requset相关的head设置

**`Requset.Header`**保存请求的head部分信息

```go
type Requset struct{
    Header Header
}

type Header map[string][]string
```
可以看到Requset将head信息保存在一个map中，我们来设置这个map即可，下面来构造一个请求head
期望服务器返回 **text/html** 类型的数据
期望服务器返回的编码类型为中文  **zh-CN,zh;**

```go
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
```
运行程序,将输出:
User-Agent=Go-http-client/1.1
Accept=text/html,q=0.8
Accept-Language=zh-CN,zh;q=0.9
Accept-Encoding=gzip

由于 **Request** 为自己构造的,所以需要再构造一个 **http.Client** ,再通过 **Do(reqeust)** 方法将我们自己构造的 **request** 发送出去.

**NewRequest()** 为http包提供一种快速构造request结构的方法，实际上这个方法的内部也是再设置Request结构的属性值，于是下面的方式与上述方式是相同的

```
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
```
上述用到了 **net/url**库中的Parse()方法，用来检测一个字符串是否是正确的url地址，返回 **url.URL** 类型与 **error**类型，除了检测以外，还可以将字符串url转换为 **url.URL** ，而 **Requst.URL**接收的是 **url.URL**类型。

**url.URL**结构还包含了其他几种方法，后面使用到的时候会再次讲解


并不是所有包含在http请求头中的信息都需要通过设置 **request.Header** 来实现，部分属性可以通过直接设 置**request.Header**的相关属性值来设置，如:
```go
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
```
开始给出的例子中Host等信息也是包含在请求头中的,http库将其单独设置为一个属性,读取和设置也就更加便利