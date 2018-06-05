> 在其他语言诸如C++/Java中, map一般都以库的方式提供 如C++中的 std::map<> ,Java中的Hashmap<>,在golang中map为内置类型,不需要引用任何库.

map是一堆未排序键值对的集合，map的key必须为可比较类型,比如 **==** 或 **!=**，map查找比线性查找快,但慢于索引查找(数组，切片)

### 定义一个map
格式  `var name map[keytype]valuetype`

其中的keytype为map中键的类型,valuetype为map中值的类型,map还可以使用 := 或者 make()创建,如:

```go
var a map[int]int
var b = map[string]string{}
c := map[string]bool{}
d := make(map[string]int)
```

###  map的初始化
map类似于数组和切片,可以在定义时直接指定初始值 如:

```go
var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
```

### map元素访问
上边定义并初始化了一个map,那么如何访问一个map中的元素呢？

可以通过键(key)来访问指定的value 如:
```go
var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
fmt.Println(a["name"],a["age"])
```
上述程序将输出: zhangsan 16



将上述的代码修改一下
```go
var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
fmt.Println(a["name"], a["age"], a["parent"])
```
map中并没有parent这个键,在诸如C++/Java或其他语言中,如果访问一个并不存在的键,将直接导致程序异常,所以判断一个map中是否存在某个键,就成了必需的步骤，但在golang中，这些都不是问题

上述代码不会发生任何异常或者产生警告
在golang中 如果访问一个未定义的key 将返回这个map中value的默认值,来验证一下

```go
package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
	fmt.Println(a["parent"])
	if a["parent"] == "" {
		fmt.Println("yes")
	}

	var b = map[int]int{1: 1, 2: 2}
	fmt.Println(b[3])
	if b[3] == 0 {
		fmt.Println("yes")
	}
}
```
上述程序将输出
//这个是空字符串
yes
0
yes

但显然我们并不想要这个结果,golang中提供了一种方式来判断数组中是否含有某个键

格式 `value,bool := map[key]` 如下列程序
```go
func main() {
	var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
	if value, ok := a["parent"]; ok == true {
		fmt.Println(value)
	} else {
		fmt.Println("key not in map")
	}
}
```
注意:golang中过的if语句中定义的变量在语句块外是访问不到的


### 修改map中的元素
和数组一样,可以使用key来修改map中的元素

```go
	var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
	fmt.Println(a)
	a["name"] = "李四"
	fmt.Println(a)
```
上述代码将输出
map[name:zhangsan age:16 sex:男]
map[name:李四 age:16 sex:男]

### 新增元素
与slice不同,map元素的新增不需要使用copy()函数 可以直接  `map[key] = value` 的方式增加元素 如:

```go
package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
	a["girlfriend"] = "苍老师"
	fmt.Println(a)
}
```

### 删除元素
比slice相比,map提供delete()函数进行元素的删除
格式 `delete(map,key)`  如下例:

```go
package main

import (
	"fmt"
)

func main() {
	var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
	delete(a, "name")
	fmt.Println(a)
}
```
以上程序将map[age:16 sex:男]

### map是引用类型

map是引用类型,来看一个简单的例子
```go
func main() {
	var a = map[int]int{1: 1, 2: 2}
	b := a
	a[1] = 123456
	fmt.Println(a, b)
}
```
以上程序会输出 map[1:123456 2:2] map[1:123456 2:2]

我们再传递进函数里面测试下,关于函数的定义后面我们会讲,此处只是一个demo

```go
func main() {
	var a = map[int]int{1: 1, 2: 2}
	b := a
	test(a)
	fmt.Println(a, b)
}

func test(a map[int]int) {
	a[1] = 111
}
```
以上程序会输出 map[2:2 1:111] map[1:111 2:2] 或者 map[1:111 2:2] map[1:111 2:2]

为什么呢?仔细思考一下,这是否再一次说明了map是无序的呢？

### map的遍历
与数组类似  map也可以使用 for  range关键字来进行遍历

```go
var a = map[string]string{"name": "zhangsan", "age": "16", "sex": "男"}
for key, value := range a {
	fmt.Println(key, value)
}
```
