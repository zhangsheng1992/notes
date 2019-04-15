> 本篇记录常用类型转换已供备忘

## 基本类型转换

### int转string

1.利用 **fmt** 标准库的 **Sprintf()** 、 **Sprintln()** 、 **Sprint()** 函数来转换

```go
var a int = 123
str := fmt.Sprintf("%d", a)
fmt.Println(str)
```

2.利用**strconv**库的 **Itoa()** 来转换

```go
var a int = 123
str := strconv.Itoa(a)
fmt.Println(str)
```
### string转int

利用**strconv**库的 **Atoi()** 函数来转换

```go
var a string = "123"
n, err := strconv.Atoi(a)
fmt.Println(n, err, r.TypeOf(n))
```
### int转int32 int64 uint等

通过类型转换即可

```go
var a int = 123
b := int32(a)
c := int64(a)
```
### int转float32 float64

```go
var a int = 123
b := float32(a)
c := float64(a)	
```

### float32 float64转int

```go
var a float32 = 1.1
b := int(a)
```

### string转[]byte

```go
var str string = "test"
var data []byte = []byte(str)
```

### []string转[]byte

需要 **encoding/gob** 与 **bytes** 库
```go
var a = []string{"a", "b", "c", "d"}
buffer := &bytes.Buffer{}
gob.NewEncoder(buffer).Encode(a)
byteSlice := buffer.Bytes()
```

### []byte转[]string
```go
buffer := &bytes.Buffer{}
backToStringSlice := []string{}
gob.NewDecoder(buffer).Decode(&backToStringSlice)
fmt.Println(backToStringSlice)
```

或

```go
b := []byte{}
str := fmt.Sprintf("%x",b)
```

### []slice转string
```go
s := []string{"a","b","c"}
str := strings.Join(s,"")
```

## 时间类型转换

### 时间戳转日期

```go
func main() {
	var timestamp int64 = 1526374941
	date := time.Unix(timep, 0).Format("06-01-02 15:04:05")
	fmt.Println(date)
}
```

### 获取当前时间戳
```go
func main(){
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)
}
```


## 常用接口类型转换

### io.ReadCloser转[]byte

常见于解析 **http.Request.Body** 与 ***http.Response.Body**  

利用ioutil标准库类转换

```go
func main(){
    //response.Body为io.ReadCloser类型
    dataByte, err2 := ioutil.ReadAll(response.Body)
    fmt.Println(dataByte, err2)	
}
```

### []byte转io.ReadCloser 

常用于设置 **http.Request.Body** 与 ***http.Response.Body** 的	内容

```go
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	str := "abc"
	request := http.Request{}
	//strings.Reader 也实现了io.ReadCloser接口 所以也可以使用
	//reader := strings.NewReader(str)
	reader := bytes.NewReader([]byte(str))
	body := ioutil.NopCloser(reader)
	request.Body = body
	fmt.Println(request)
}
```