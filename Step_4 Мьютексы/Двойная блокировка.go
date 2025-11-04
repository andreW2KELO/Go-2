package main

import "sync"

var m sync.Mutex

func first() {
	m.Lock()
	defer m.Unlock()
	second() // lock
}

func second() {
	m.Lock() // здесь будет вызов m.Lock() второй раз
	defer m.Unlock()
	// далее — основное тело функции
}

func main() {
	go first()
	go second()
}
