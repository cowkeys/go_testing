package lib

import "testing"

func TestDelete(t *testing.T) {
	Delete()
}

func Benchmark(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Delete()
	}
}
//目录下 执行
//go test
// go test -test.bench=".*"
