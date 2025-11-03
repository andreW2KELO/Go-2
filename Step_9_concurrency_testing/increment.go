package main

import "sync"

var mu sync.Mutex
var counter int

func Increment() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func GetCounter() int {
	mu.Lock()
	defer mu.Unlock()
	return counter
}
