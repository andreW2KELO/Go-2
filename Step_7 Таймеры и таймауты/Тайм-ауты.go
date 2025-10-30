package main

import (
	"fmt"
	"time"
)

func timeoutTest() string {
	// уменьшаем время здесь, чтобы попасть в пределы
	time.Sleep(4 * time.Second)
	return "Функция TimeoutTest выполнена!"
}

func main() {
	// создание канала C
	c := make(chan string, 1)
	// создание потока выполнения горутины
	go func() {
		str := timeoutTest()
		c <- str
	}()
	// создание тайм-аута для выполнения функции
	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Время вышло!")
	}
}
