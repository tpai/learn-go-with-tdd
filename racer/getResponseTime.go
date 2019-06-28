package main

import (
	"fmt"
	"net/http"
	"time"
)

func measureResponseTime(url string) chan time.Duration {
	ch := make(chan time.Duration)
	go func() {
		startTime := time.Now()
		http.Get(url)
		ch <- time.Since(startTime)
	}()
	return ch
}

func main() {
	urls := []string{"https://google.com", "https://facebook.com"}
	for _, url := range urls {
		timeCost := <-measureResponseTime(url)
		fmt.Printf("%s %v\n", url, timeCost)
	}
}
