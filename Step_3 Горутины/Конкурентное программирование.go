package main

import (
	"fmt"
	"time"
)

func doSomething() {
	fmt.Println("hello world")
}

func main() {
	go doSomething()
	time.Sleep(1 * time.Second)
	// В этом примере функция doSomething() запускается в горутине, затем вызывается time.Sleep(),
	// чтобы основная горутина не завершилась до того, как doSomething() завершит свою работу.
	// Благодаря time.Sleep() программа успевает вывести "Hello, World!" перед тем, как завершится.
}
