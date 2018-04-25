>作为正式学习go语言的第一步,我们应该尊重传统,所以本节我们来完成一个小程序

### 编写第一个go程序
在桌面新建一个记事本,重命名为**hello.go**文件,然后输入如下内容

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println("hello world!")
}
```

### 打开终端并运行
**windows+R** 然后输入 **cmd** ,输入下列内容并回车

```shell
cd desktop
go run hello.go
```

如果你可以看到
hello world!
那么恭喜你,你已经迈出了伟大的第一步

(ps:如果提示找不到文件,请自行百度学习一下相对路径与绝对路径,这个知识在所有语言中都是通用的,而且非常容易理解)

### 注意

自己阅读上述代码,go语言的代码有如下要求

go程序文件必须以 ** .go **结尾,不要用特殊符号命名,尽量不要使用数字开头的文件名

每个go的文件都必须包含** package **的声明,用来说明他属于哪个包的,一个包可以包含多个go文件

我们后面学习的大部分时候,这里是固定的,即 package main

** import ** 用来导入库,库就是官方或者他人写好的代码,导出以后就可以使用了,在这里我门引入一个库 **fmt**

fmt就是官方的一个库,用来输出,在基础部分,我们仅需要记住用法 

```go
package main//package用来指定代码所属代码库(也可以叫包,为了避免冲突,后面统一称为库)

import "fmt"//import用来引入其他库文件

func main(){
    fmt.Println("你想要输出的东西")
}

```

如果一次性用到多个库文件,可以使用如下方式倒入,每行一个,必须用双引号括起来
```go
import (
    "fmt"
    "os"
    "io"
)

```

** main() ** 如果你有C/C++基础,那么应该知道程序运行的时候需要一个主函数,或者说叫入口函数,与C/C++类似,go语言中的主函数也叫main()

main(){}是必须的,而且只能有一个,且不能有参数或返回值

### go程序的运行顺序

无论你有多少个go代码文件,倒入了多少个其他库,甚至文件的名称是什么,go的运行都是固定的

**go的程序总是从package main 中的 main()开始执行**,并按照自上而下的顺序(注)依次执行完所有代码后退出main()函数

退出main()函数也意味着程序运行完成,不如下面这个例子

```go
package main

import "fmt"

func main(){
    fmt.Println(1)
    fmt.Println(true)
    fmt.Println("a")
}

```
他会从上到下输出1 true 和 a,在输出a以后,由于没有后续需要执行的代码,main()函数也就执行完毕并退出,程序停止运行


注:这里想说的是不考虑程序运行出错,协程,goto等语句

### go fmt大法

go语言彻底解决了一个困扰所有程序猿数十年,并因此爆发过多场讨论与战争的问题,那就是代码的书写格式(python笑而不语)

当你写完一个go文件以后,在终端下运行 go fmt 文件名,比如,


```shell
go fmt hello.go

```

这个工具会自动的将你的go文件转换为同一风格,熟悉这种风格,当你阅读别人的代码的时候,你会发现,所有人的程序看上去都是那么亲切

ps:其实并没有战争的啦,下面给出两个链接,可以当乐子看看,go fmt大法好

[花括号圣战](http://mp.weixin.qq.com/s?__biz=MzAxMzMxNDIyOA==&mid=215114843&idx=1&sn=5a765de3c9a0ab60ebe193eee09770f9 "花括号圣战")

[缩进圣战](http://mp.weixin.qq.com/s?__biz=MzAxMzMxNDIyOA==&mid=215766844&idx=1&sn=59b52569b7c52ac874e0b2fdb4bce3f1&scene=0#rd "缩进圣战")


### go build

go语言是跨平台的,一次编写在所有支持的平台上都可以运行,而且不依赖于宿主,但编写后的go文件不能直接运行(注),需要先编译

go build命令就是用来编译go文件,通过编译参数不同,将同一份代码编译成可以在各自平台运行的版本,用法

```shell
go build 文件名
```

windows下会编译成**文件名.exe**的可执行文件,macos X,linunx下会编译成执行文件,都可以直接运行

注:go run 命令实际上是先执行go build命令,然后在执行编译后的可执行文件的

go编译起会将所有用到的库都编译进去,不会重复引用,很智能,下面这张图代表了编译时的顺序,这个可以等学到后边再回头来理解

<img src="/public/images/wiki/build.png" class="blogimage" />