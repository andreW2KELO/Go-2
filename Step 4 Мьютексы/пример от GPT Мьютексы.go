package main

import (
	"fmt"
	"time"
)

var counter int

func increment() {
	for i := 0; i < 1000; i++ {
		counter++ // несколько горутин меняют одну переменную
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go increment()
	}
	time.Sleep(time.Second)
	fmt.Println("Counter:", counter)
}
