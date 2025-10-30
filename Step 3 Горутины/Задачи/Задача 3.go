package main

import "fmt"

func Send5(ch1, ch2 chan int) {
	go func() {
		ch1 <- 0
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		ch2 <- 0
		ch2 <- 1
		ch2 <- 2
	}()
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	Send5(ch1, ch2)
	fmt.Println(<-ch1)
}
