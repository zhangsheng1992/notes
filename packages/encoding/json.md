>json是一种轻量级的数据交互格式,在日常开发中,json常被用来作为两个系统间交互的格式,本章来学习go语言是如何处理 json的

### encoding/json

go语言内置了解析json的标准库 **encoding/json** (下面简称为json库)来生存与解析json数据,json库比较庞大,此处仅讲解日常常用的部分,足以涵盖实际应用场景

### 解析字符串json

首先来看一个简单的json字符串

```json
{"name":"zhangsan","age":16}
```

是不是很像 **map** 结构?
在实际开发时,比如我们提供一个web API 接口供app提交数据,无论是ios平台的object-c/swift,或者安卓平台的java语言,与go语言的数据类型都是不一样的,所以需要一个中间人,而json就是充当这个中间人,一次简单的数据传递可以表示如下:

用户在app上输入帐号密码 => app语言将帐号密码封装为json格式 =>发送到服务端 =>服务端解析json数据 => 得到用户的帐号密码

如何来解析上面的json字符串？既然很像map,那么就将它解析称为一个map，json库解析json字符串使用 **Unmarshal()** 方法来实现,函数签名如下:

```go
	func Unmarshal(data []byte, v interface{}) error {}
```
Unmarshal会将[]byte类型的json数据转换到一个interface{}中,所以需要传递指针参数,下面来看如何使用( **\** 是go中的转移字符 并非json数据格式,下同)

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
var a string = "{\"name\":\"zhangsan\",\"age\":16}"
	var data = map[string]string{}

	err := json.Unmarshal([]byte(a), &data)
	fmt.Println(data,err)
}
```
程序输出结果 map[name:zhangsan] json: cannot unmarshal number into Go value of type string

仔细观察一下,为什么age没有解析出来呢？ 原因是我门定义的age是一个 **int** 类型,而定义的map是一个 **map[string]string** 类型

### json数据断言

json数据类型常用的有四种, 字符,整形,浮点和布尔类型,比如下面这个json数据,咋办？

```go
	var a string = "{\"name\":\"zhangsan\",\"age\":16,\"man\":true}"
```

可以使用interface{},然后再通过断言

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a string = "{\"name\":\"zhangsan\",\"age\":\"16\",\"man\":true}"
	var data = map[string]interface{}{}
	
	err := json.Unmarshal([]byte(a), &data)
	fmt.Println(data, err)

	for k, v := range data {
		switch v.(type) {
		case string:
			fmt.Println(k, "的值是string 类型,应该是name字段")
		case int:
			fmt.Println(k, "的值是int 类型,应该是age字段")
		case bool:
			fmt.Println(k, "的值是bool 类型,应该是man字段")
		}
	}
}
```
看,成功解析了这个串,只是在使用的时候需要断言一下

### 解析复杂的json数据

实际中我们要处理的json绝不可能这么简单,来看一下稍微贴近实际点的json数据:

```json
{
	"user":{
		"name":"zhangsan",
		"level":99,
		"other":{
			"money":"11$",
			"Vip":true
		}
	},
	"team":{
		"teamname":"xxx",
		"member":[
			{
				"tname":"lisi"
			},
			{
				"tname":"wangwu"
			}
		]
	}
}
```

上述数据描述了一个名为zhangsan的用户的信息和他的工作组信息,对这种结构的数据,如果使用断言，你得写多少？有没有一种简单的办法呢？

### 从json文件中读取json数据

针对上面的问题,可以使用结构体来解析这个数据,在此再介绍json库的一个结构 Decoder,它允许从输入源中读取并解析json数据,把上面的json数据保存进一个test.json文件中,然后来解析它

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	var v = map[string]interface{}{}
	decoder.Decode(&v)
	fmt.Println(v)
}
```
利用interface{}的方式成功解析,下面来讲解如何利用结构体解析

定义一个结构体Student,其中包含两个结构体成员User和Team,每个成员属性若干

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	User User
	Team Team
}

type User struct {
	Name  string
	Level int
	Other Other
}

type Other struct {
	Money string
	Vip   bool
}

type Team struct {
	Teamname string
	Member   []map[string]string
}

func main() {
	file, err := os.Open("test.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	var v = Student{}
	err1 := decoder.Decode(&v)
	fmt.Println(v, err1)
}
```

{{zhangsan 99 {11$ true}} {xxx [map[tname:lisi] map[tname:wangwu]]}}

### 使用结构体解析json时的注意事项

1.蛋疼的大小写,因为结构体的访问性(小写的字段其它包访问不了,这里的其它包指的就是json包),所以要解析的字段首字母都必须大写

2.结构体属性名称需与json中的字段一致(首字母的大写在解析时会自动转换成小写)

3.如果你的结构体中属性名称与json中的不一致则解析不了,可以通过指的别名来处理
```go
type Other struct {
	Qian string `json:"money"`
	Vip  bool
}
```
用Qian这个属性来接受money字段,在后边需要用 `json:"目标字段"`来表明
蛋疼归蛋疼,但是标准库就是这么弄的,希望有更好的官方库或者三方库来完善


### 转换为json

将go数据类型转换为json相比较下就比较简单了 json库使用 **Marshal()** 方法来生成json数据

**将一个map转换为json（字符串 数组 interface{}等类型类似,传进函数就可以,有返回值,不需要传递指针）**

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a = map[string]interface{}{"name": "zhangsan", "age": 25}
	data, err := json.Marshal(a)
	fmt.Println(string(data), err)
}
```
{"age":"lisi","name":"zhangsan"}

**将一个结构体转换为json（蛋疼的大小写与字段名依旧存在）**

```go
package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var a = A{"zhangsan", 25}
	data, err := json.Marshal(a)
	fmt.Println(string(data), err)
}
```
如果不制定别名,转换成的json字段名为大写的

有 **Decoder** 结构那么也就有 **Encoder** 结构,篇幅原因在此不做过多叙述,有兴趣的可以自行学习一下










