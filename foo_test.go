package foo

import (
	"testing"
)

func BenchmarkPrintFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processFoo()
	}
}

func BenchmarkPrintFooPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processFooPointer()
	}
}
