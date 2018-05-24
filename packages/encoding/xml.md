>xml是一种可扩展的标记语言,被设计用来传输和存储数据,本章讲解如何使用go来解析与生成xml

### xml的特点

1.大小写敏感
2.结构简单
3.标签式的写法
在实际使用中,xml作为则多一种数据模版,如
- 安卓开发的界面构建
- java web开发中的配置文件
- 微信图文素材的模版

###与json相比
json作为一种简单通信的数据格式,由于易用性所以被广泛使用,但是json有如下不足
1.结构化 json是一种键值对的数据格式,所以没有结构化的概念,即：

```json
	{"a":"1","b":"2"}
```
与

```json
	{"b":"2","a":"1"}
```
是完全相同的

2.不支持节点属性 有时候为了实现节点属性会使用嵌套结构代替,这会使得阅读非常困难，如:

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

### 一个简单的xml

下面就是一个简单的xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<school>
	<a id="1">
		<name>zhangsan</name>
		<age>25</age>
	</a>
	<a id="2">
		<name>lisi</name>
		<age>26</age>
	</a>
</school>
```

有如下特点
1.xml由一个一个节点组成
2.所有节点都有一个根节点
3.节点可以包含任意个属性与值

### 解析xml

以上面的xml为例,现在来进行解析,xml不同于json的键值对结构,所以无法直接将简单的xml转换为json,相比json,他的转换复杂一点,将上边的xml保存进test.xml中

```go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type School struct {
	Student []Student `xml:"a"`
}

type Student struct {
	Id   string `xml:"id,attr"`
	Name string `xml:"name"`
	Age  int    `xml:"age"`
	Content string `xml:",innerxml"`
}

func main() {
	data, _ := ioutil.ReadFile("test.xml")
	var a = School{}
	err := xml.Unmarshal(data, &a)
	fmt.Println(a, err)
}
```

xml解析与生产使用 **encoding/xml** 库来完成,后面简称为xml,函数与json完全相同，但在定义结构体的时候需要注意如下

1.结构和属性的首字母大写(go的访问控制缘故)
2.每个属性或者节点都需要指明类型

`xml:"id,attr"` 表示是属性节点  id为属性的名称
`xml:"name"` 表明是子节点  name为xml的节点名称
`xml:",innerxml"` 表示取当前节点的所有子节点内容

**xml** 库使用 **Unmarshal()** 来完成解析,签名如下:

```go
	func Unmarshal(data []byte, v interface{}) error {}
```

是不是和json很像呢？对,非常像,上面我们是使用 **ioutil.ReadFile()** 读取的xml byte来解析,xml库也有Decoder结构,可以让我们直接读取一个xml文件并解析,稍微该一下上面的例子
```go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type School struct {
	Student []Student `xml:"a"`
}

type Student struct {
	Id      string `xml:"id,attr"`
	Name    string `xml:"name"`
	Age     int    `xml:"age"`
	Content string `xml:",innerxml"`
}

func main() {
	file, _ := os.Open("test.xml")
	decoder := xml.NewDecoder(file)
	var a = School{}
	err := decoder.Decode(&a)
	fmt.Println(a, err)
}
```

### xml的生成

生产xml使用 **Marshal()** 方法来完成,函数签名如下
```go
func Marshal(v interface{}) ([]byte, error) {}
```
举个例子:
```go
package main

import (
	"encoding/xml"
	"fmt"
)

type School struct {
	Student []Student
}

type Student struct {
	Name string
	Age  int
}

func main() {
	var zhangsan = Student{"张三", 15}
	var lisi = Student{"李四", 25}
	var a = School{[]Student{zhangsan, lisi}}
	data, err := xml.Marshal(a)
	fmt.Println(string(data), err)
}
```

会生成如下xml

```xml
<School>
	<Student>
		<Name>张三</Name>
		<Age>15</Age>
	</Student>
	<Student>
		<Name>李四</Name>
		<Age>25</Age>
	</Student>
</School>
```

与解析一样,可以使用tag标签来指定属性

```go
package main

import (
	"encoding/xml"
	"fmt"
)

type School struct {
	Student []Student
}

type Student struct {
	Id   int `xml:"id,attr"`
	Name string
	Age  int
}

func main() {
	var zhangsan = Student{1, "张三", 15}
	var lisi = Student{2, "李四", 25}
	var a = School{[]Student{zhangsan, lisi}}
	data, err := xml.Marshal(a)
	fmt.Println(string(data), err)
}
```

```xml
<School>
	<Student id="1">
		<Name>张三</Name>
		<Age>15</Age>
	</Student>
	<Student id="2">
		<Name>李四</Name>
		<Age>25</Age>
	</Student>
</School>
```