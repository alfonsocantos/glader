package ttl

import "time"

func eraser(g *Config) {

	for {

		l := g.List()
		now := time.Now()

		for _, id := range l {

			t := getTTL(g, id)
			if now.After(t) {
				g.Delete(id)
			}
		}

		time.Sleep(1 * time.Second)
	}

}

func getTTL(gl *Config, id string) time.Time {

	gl.lock.Lock()
	defer gl.lock.Unlock()

	return gl.ttlMap[id]
}
