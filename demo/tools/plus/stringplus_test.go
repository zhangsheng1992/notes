package plus_test

import (
	p "plus"
	"testing"
)

func BenchmarkTestPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Test1()
	}
}

func BenchmarkTestBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Test2()
	}
}

func BenchmarkTestFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Test3()
	}
}

func BenchmarkTestAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Test4()
	}
}

func BenchmarkTestAppendByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.Test5()
	}
}

func CapTest(T *testing.T) {
	p.Test6()
}
