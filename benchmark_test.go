package glader

import (
	"github.com/alfonsocantos/glader/memory"
	"testing"
)

func BenchmarkGlader(b *testing.B) {
	gl := memory.New()
	for i := 0; i < b.N; i++ {
		gl.Add("my-item", struct{ key1 string }{key1: "value1"})
		gl.List()
		gl.Get("my-item")
		gl.Delete("my-item")
	}
}

func BenchmarkGet(b *testing.B) {
	gl := memory.New()
	gl.Add("my-item", struct{ key1 string }{key1: "value1"})
	for i := 0; i < b.N; i++ {
		gl.Get("my-item")
	}
}

func BenchmarkList(b *testing.B) {
	gl := memory.New()
	gl.Add("my-item", struct{ key1 string }{key1: "value1"})
	for i := 0; i < b.N; i++ {
		gl.List()
	}
}

func BenchmarkAdd(b *testing.B) {
	gl := memory.New()
	for i := 0; i < b.N; i++ {
		gl.Add("my-item", struct{ key1 string }{key1: "value1"})
	}
}
