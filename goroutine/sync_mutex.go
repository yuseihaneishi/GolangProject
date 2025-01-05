package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	defer c.mux.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	var wg sync.WaitGroup
	c := Counter{v: make(map[string]int)}

	wg.Add(2) // 2つのゴルーチンを待機

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c.Inc("key")
		}
	}()

	wg.Wait() // すべてのゴルーチンが終了するまで待機
	fmt.Println(c, c.Value("key"))
}
