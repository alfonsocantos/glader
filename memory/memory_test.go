package memory

import (
	"context"
	"github.com/alfonsocantos/glader/ttl"
	"testing"
	"time"

	. "github.com/fulldump/biff"
)

func environment(f ...func(ctx context.Context)) {

	ctx := context.Background()

	gl := New()

	ctx = context.WithValue(ctx, "mem-test", gl)

	for _, f := range f {
		f(ctx)
	}
}

func getGlader(ctx context.Context) *Glader {
	return ctx.Value("mem-test").(*Glader)
}

func TestGlader_Add(t *testing.T) {

	environment(
		func(ctx context.Context) {
			gl := getGlader(ctx)

			gl.Add("my-item", struct{ key1 string }{key1: "value1"})
		},
		func(ctx context.Context) {
			gl := getGlader(ctx)

			AssertEqual(gl.mem["my-item"], struct{ key1 string }{key1: "value1"})

		})
}

func TestGlader_Get(t *testing.T) {

	environment(
		func(ctx context.Context) {
			gl := getGlader(ctx)

			gl.mem["my-item"] = struct{ key1 string }{key1: "value1"}
		},
		func(ctx context.Context) {
			gl := getGlader(ctx)

			item := gl.Get("my-item")
			AssertEqual(item, struct{ key1 string }{key1: "value1"})

		})
}

func TestGlader(t *testing.T) {

	environment(
		func(ctx context.Context) {
			gl := getGlader(ctx)

			gl.Add("my-item", struct{ key1 string }{key1: "value1"})
		},
		func(ctx context.Context) {
			gl := getGlader(ctx)

			item := gl.Get("my-item")
			AssertEqual(item, struct{ key1 string }{key1: "value1"})

		},
		func(ctx context.Context) {
			gl := getGlader(ctx)

			err := gl.Delete("my-item")
			AssertNil(err)
		},
		func(ctx context.Context) {
			gl := getGlader(ctx)

			item := gl.Get("my-item")
			AssertEqual(item, nil)

		})
}

func TestGlader_AddWithTTL(t *testing.T) {

	environment(
		func(ctx context.Context) {
			gl := getGlader(ctx)
			ttlGlader := ttl.New(gl, time.Second)

			ttlGlader.Add("my-item", struct{ key1 string }{key1: "value1"})
			AssertNotNil(gl.Get("my-item"))
			time.Sleep(2 * time.Second)
			AssertNil(gl.Get("my-item"))

		})
}
