package memory

import (
	"sync"
)

type Glader struct {
	mem map[string]interface{}

	lock sync.RWMutex
}

func New() *Glader {

	g := &Glader{
		mem: map[string]interface{}{},
	}

	return g
}

func (g *Glader) Get(id string) interface{} {

	g.lock.Lock()
	defer g.lock.Unlock()

	return g.mem[id]
}

func (g *Glader) List() []string {

	g.lock.Lock()
	defer g.lock.Unlock()

	l := make([]string, len(g.mem))

	i := 0
	for key := range g.mem {
		l[i] = key
	}

	return l
}

func (g *Glader) Add(id string, item interface{}) {

	g.lock.Lock()
	defer g.lock.Unlock()

	g.mem[id] = item

	return
}

func (g *Glader) Delete(id string) error {

	g.lock.Lock()
	defer g.lock.Unlock()

	delete(g.mem, id)

	return nil
}
