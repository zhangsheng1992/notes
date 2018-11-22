package plus

import (
	"bytes"
	"fmt"
)

var a string = "Hello,"
var b string = "World!"

var c = []byte{72, 101, 108, 108, 111, 44}
var d = []byte{87, 111, 114, 108, 100, 33}

var cArr = [12]byte{72, 101, 108, 108, 111, 44}

// + 方法拼接
func Test1() string {
	return a + b
}

// buffer拼接
func Test2() string {
	var buffer bytes.Buffer
	buffer.WriteString(a)
	buffer.WriteString(b)
	return buffer.String()
}

// format拼接
func Test3() string {
	return fmt.Sprint(a, b)
}

// append方式
func Test4() string {
	return string(append([]byte(a), []byte(b)...))
}

// 不进行类型转换
func Test5() []byte {
	return append(c, d...)
}

// 长度容量
func Test6() {
	e := []byte(a)
	f := []byte(b)
	fmt.Println("c的长度:", len(c), "容量:", cap(c))
	fmt.Println("d的长度:", len(d), "容量:", cap(d))
	fmt.Println("e的长度:", len(e), "容量:", cap(e))
	fmt.Println("f的长度:", len(f), "容量:", cap(f))
}
