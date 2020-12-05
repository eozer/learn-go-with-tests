package mysync

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex
	// NOTE: we can also directly use this -> sync.Mutex as c.Lock() or c.Unlock()
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
