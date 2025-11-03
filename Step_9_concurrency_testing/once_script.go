package main

import (
	"sync"
)

var (
	once     sync.Once
	counter  int
	incrFunc func()
)

// initialize увеличивает счётчик только один раз
func initialize() {
	counter++
}

// Increment вызывает initialize только один раз
func Increment() {
	once.Do(initialize)
}

// GetCounter возвращает текущее значение счётчика
func GetCounter() int {
	return counter
}
