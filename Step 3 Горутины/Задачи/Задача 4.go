package main

func Process(nums []int) chan int {
	ch := make(chan int, 10)
	for _, n := range nums {
		ch <- n
	}
	close(ch)
	return ch
}
