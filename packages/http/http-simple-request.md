>http request是http库内置结构,除了处理http请求,也可以用来发送http请求

http package为 ** Request ** 结构体封装好了发送请求的方法,对用户而言,发送一个http请求,实际上就是设置Request结构中的相关属性,下面将逐一讲解各个属性的含义


### Ruquest结构

Request结构如下:

```go
type Request struct {
	Method string
	URL *url.URL
	Proto      string
	ProtoMajor int
	ProtoMinor int
	Header Header
	Body io.ReadCloser
	GetBody func() (io.ReadCloser, error)
	ContentLength int64
	TransferEncoding []string
	Close bool
	Host string
	Form url.Values
	PostForm url.Values
	MultipartForm *multipart.Form
	Trailer Header
	RemoteAddr string
	RequestURI string
	TLS *tls.ConnectionState
	Cancel <-chan struct{}
	Response *Response
	ctx context.Context
}
```

http库提供了Get()方法来发送GET请求,来看一个最简单的GET请求,server与示例代码请提前下载 地址 [server.go](https://github.com/zhangsheng1992/notes/blob/master/demo/http-client/simple/server.go)

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
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
```

** `go run server.go` ** 启动demo中的http server
然后** `go run client.go` ** 
运行将会输出:

a = 1
b = 2
c = abcd

来看一下http.Get()方法,下面已整理好关键代码
```
func Get(url string) (resp *Response, err error) {
	return DefaultClient.Get(url)
}

var DefaultClient = &Client{}

func (c *Client) Get(url string) (resp *Response, err error) {
	req, err := NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func NewRequest(method, url string, body io.Reader) (*Request, error) {
	if method == "" {
		method = "GET"
	}
	if !validMethod(method) {
		return nil, fmt.Errorf("net/http: invalid method %q", method)
	}
	u, err := parseURL(url)
	if err != nil {
		return nil, err
	}
	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	// The host's colon:port should be normalized. See Issue 14836.
	u.Host = removeEmptyPort(u.Host)
	req := &Request{
		Method:     method,
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(Header),
		Body:       rc,
		Host:       u.Host,
	}
	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			req.ContentLength = int64(v.Len())
			buf := v.Bytes()
			req.GetBody = func() (io.ReadCloser, error) {
				r := bytes.NewReader(buf)
				return ioutil.NopCloser(r), nil
			}
		case *bytes.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return ioutil.NopCloser(&r), nil
			}
		case *strings.Reader:
			req.ContentLength = int64(v.Len())
			snapshot := *v
			req.GetBody = func() (io.ReadCloser, error) {
				r := snapshot
				return ioutil.NopCloser(&r), nil
			}
		default:
		}
		if req.GetBody != nil && req.ContentLength == 0 {
			req.Body = NoBody
			req.GetBody = func() (io.ReadCloser, error) { return NoBody, nil }
		}
	}

	return req, nil
}
```

** NewRequest() ** 方法中有一段 ** `req := &Request{}` ** 

实质上GET()方法就是简单的设置了 ** Requset ** 结构体的 ** Method ** 属性与 ** URL ** 属性,然后使用 ** Do() ** 方法来发送这个请求，返回一个 ** Request ** 结构的 ** Response ** 属性,其中保存着请求的结果.

** Do() ** 方法与 ** Response **结构会在后边讲解.

http库还提供了另一种方法 Post(),是否可以猜测Post()方法也是如此？



```go
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

```

** `go run client.go` ** 运行结果:
the request body is a=1&b=2









