package main

import "sync"

var (
	Buf   []int
	mutex sync.Mutex
)

func Write(num int) {
	mutex.Lock()
	defer mutex.Unlock()
	Buf = append(Buf, num)
}

func Consume() int {
	mutex.Lock()
	defer mutex.Unlock()
	if len(Buf) == 0 {
		return 0
	}
	res := Buf[0]
	Buf = Buf[1:]
	return res
}

func main() {
	// Пример использования
	Write(10)
	Write(20)
	Write(30)

	println(Consume())
	println(Consume())
	println(Consume())
	Consume() // Пустой буфер
}
