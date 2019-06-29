package main

import (
	"fmt"
	"sync"
)

type LockCounter struct {
	mu    sync.Mutex
	value int
}

func (c *LockCounter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *LockCounter) Get() int {
	return c.value
}

type ChannelCounter struct {
	value int
}

func (c *ChannelCounter) Add() {
	c.value++
}

func (c *ChannelCounter) Get() int {
	return c.value
}

func main() {
	wantedCount := 100

	lcounter := LockCounter{}
	var wg sync.WaitGroup
	wg.Add(wantedCount)

	for i := 0; i < wantedCount; i++ {
		go func(w *sync.WaitGroup) {
			lcounter.Add()
			w.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Printf("LockCounter=%d\n", lcounter.Get())

	var result ChannelCounter
	ccounter := ChannelCounter{}
	channel := make(chan ChannelCounter)

	for i := 0; i < wantedCount; i++ {
		go func() {
			ccounter.Add()
			channel <- ccounter
		}()
		result = <-channel
	}
	fmt.Printf("ChannelCounter=%d\n", result.Get())
}
