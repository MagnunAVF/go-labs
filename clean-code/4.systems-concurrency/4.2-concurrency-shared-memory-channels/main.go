package main

// The Bad: Shared State with Mutexes While "correct," it's easy to forget to
// Unlock(), leading to deadlocks, and it's harder to track data flow.
// type Counter struct {
//     count int
//     mu    sync.Mutex
// }

// func (c *Counter) Increment() {
//     c.mu.Lock()
//     defer c.mu.Unlock()
//     c.count++
// }

// The Good: Communication via Channels Data is "owned" by one goroutine. This
// eliminates the need for locks and makes the system a "pipeline" of data.
type Counter struct {
	increments chan struct{}
	count      int
}

func NewCounter() *Counter {
	c := &Counter{
		increments: make(chan struct{}),
	}
	// Only this goroutine touches 'count'
	go func() {
		for range c.increments {
			c.count++
		}
	}()
	return c
}

func (c *Counter) Increment() {
	c.increments <- struct{}{}
}
