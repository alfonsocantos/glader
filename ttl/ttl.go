package ttl

import (
	"github.com/alfonsocantos/glader"
	"sync"
	"time"
)

type Config struct {
	ttlMap map[string]time.Time
	ttl    time.Duration

	lock sync.RWMutex

	store glader.Glader
}

func New(store glader.Glader, ttl time.Duration) *Config {
	g := &Config{
		ttlMap: map[string]time.Time{},
		ttl:    ttl,

		store: store,
	}

	go eraser(g)

	return g
}

func (g *Config) Get(id string) interface{} {

	return g.store.Get(id)
}

func (g *Config) List() []string {

	return g.store.List()
}

func (g *Config) Add(id string, item interface{}) {

	g.lock.Lock()
	defer g.lock.Unlock()

	g.store.Add(id, item)
	g.ttlMap[id] = time.Now().Add(g.ttl)

	return
}

func (g *Config) Delete(id string) error {

	g.lock.Lock()
	defer g.lock.Unlock()

	g.store.Delete(id)
	delete(g.ttlMap, id)

	return nil
}
