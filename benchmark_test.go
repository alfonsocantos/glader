package glader_test

import (
	"github.com/alfonsocantos/glader/memory"
	"github.com/alfonsocantos/glader/ttl"
	"github.com/google/uuid"
	"sync"
	"testing"
	"time"
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

func BenchmarkAddWithTTL(b *testing.B) {
	gl := memory.New()
	ttlGlader := ttl.New(gl, time.Millisecond)
	for i := 0; i < b.N; i++ {
		ttlGlader.Add("my-item", struct{ key1 string }{key1: "value1"})
	}
}

func BenchmarkCaos(b *testing.B) {
	gl := memory.New()
	ttlGlader := ttl.New(gl, time.Second)
	data := struct {
		field1 string
		field2 string
	}{
		field1: "abcdefghijklmnñopqrstuvwxyz",
		field2: "ABCDEFGHIJKLMNÑOPQRSTUVWXYZ",
	}

	for i := 0; i < b.N; i++ {
		getWg := sync.WaitGroup{}
		listWg := sync.WaitGroup{}
		addWg := sync.WaitGroup{}
		addTTLWg := sync.WaitGroup{}
		deleteWg := sync.WaitGroup{}

		// randoms gets
		go func() {
			getWg.Add(1)
			defer getWg.Done()
			for a := 0; a < 10; a++ {
				gl.Get(uuid.NewString())
			}
		}()

		// random lists
		go func() {
			listWg.Add(1)
			defer listWg.Done()
			for a := 0; a < 10; a++ {
				gl.List()
			}
		}()

		// random adds
		go func() {
			addWg.Add(1)
			defer addWg.Done()
			for a := 0; a < 10; a++ {
				gl.Add(uuid.NewString(), data)
			}
		}()

		go func() {
			addTTLWg.Add(1)
			defer addTTLWg.Done()
			for a := 0; a < 10; a++ {
				ttlGlader.Add(uuid.NewString(), data)
			}
		}()

		// random deletes
		go func() {
			deleteWg.Add(1)
			defer deleteWg.Done()
			for a := 0; a < 10; a++ {
				gl.Delete(uuid.NewString())
			}
		}()

		getWg.Wait()
		listWg.Wait()
		addWg.Wait()
		addTTLWg.Wait()
		deleteWg.Wait()
	}

}
