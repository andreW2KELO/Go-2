package main

import (
	"fmt"
	"sync"
	"time"
)

var counter2 int
var mu sync.Mutex // объявляем мьютекс

func increment2() {
	for i := 0; i < 1000; i++ {
		mu.Lock()   // блокируем доступ другим горутинам
		counter2++  // безопасно изменяем общую переменную
		mu.Unlock() // разблокируем
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go increment2()
	}
	time.Sleep(time.Second)
	fmt.Println("Counter:", counter2)
}
