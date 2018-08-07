> md5算是最常用的散列加密算法了，golang的标准库也支持md5

### crypto/md5
md5算法包含在 **crypto/md5** 下，md5算法并不复杂，因此这个库相对简单，外部可访问的方法与常量相对较少。

### 常量 BlockSize Size

BlockSize md5算法是分块计算的,512位为一个块，即64字节。

Size 定义了多少块每次校验的字节数，16字节。


### 对字符串进行md5------md5.Sum()
函数签名如下，接收byte类型参数，返回byte数组

```go
    func Sum(data []byte) [Size]byte
```

对于简单的字符串可以直接进行加密, **int** ， **float** 等类型可能需要先格式化为字符串。

```go
    str := "123456"
    re := md5.Sum([]byte(str))
    fmt.Printf("str %s 的md5为 %x", str, re)
```


### md5.New()
md5.New()返回一个hash.Hash对象，函数签名如下

```go
    func New() hash.Hash
```

hash.Hash提供了几个方便的方法

### hash.Write()

函数签名如下:

```go
    func Write(p []btye) (n int, err error)
```

Write()方法允许我们向hash对象中写入数据,md5是分片计算的,因此可以将一个较大的字符串切割成多个部分，每个部分分别计算md5，然后再整体计算，这在计算一个较长的字符串时非常有效


### hash.Sum()

Sum()方法与md5.Sum()类似

```go
	func Sun(b []btye) b []byte
```

### hash.Reset()

Reset()方法会清空Write()方法写入的数据。


### demo------计算大文件的md5
假设有一个2GB的文件，需要计算这个文件的md5，官方库没有类似的filemd5()的方法，因此需要自己计算

计算步骤
1. 打开文件
2. 读取文件内容
3. 计算md5

来看一下实现:

```go
    file := "a.mp4"
    f, _ := os.Open(file)
	// 1
    fileinfo, _ := ioutil.ReadAll(f)
    // 2
	re := md5.Sum(fileinfo)
    fmt.Printf("文件的md5为%x", re)
```

这样计算没有问题，可以计算出文件的md5，但是耗费内存太高。

文件大小为2GB，那么在 **1 ** 位置时的 **fileinfo  ** 占用了2GB内存。
在 **2 ** 的时候，由于是值类型，参数传递的时候会做拷贝，因此会再消耗2GB内存。


### 使用Write()来优化md5计算

计算步骤
1. 打开文件
2. 读取文件内容的一部分
3. 计算md5
4. 循环执行2,3步
5. 读取完毕,合并计算md5


```go
	filepath := "./a.mp4"
	file, _ := os.Open(filepath)
	buffer := make([]byte, 1024*1024*20)
	var offset int64 = 0
	m := md5.New()

	for {
		n, err := file.ReadAt(buffer, offset)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}

		if err != nil && err == io.EOF {
			m.Write(buffer[:n])
			break
		}
		
		offset = offset + int64(n)
		m.Write(buffer)
	}
	filemd5 = fmt.Sprintf("%x", m.Sum(nil))
	fmt.Printf("文件的md5为%x", filemd5)
```

这样耗费的内存就为你定义的buffer的2倍，此处为20MB。
