package main

import (
	"errors"
	"fmt"
	"time"
)

func TimeoutFibonacci(n int, timeout time.Duration) (int, error) {
	c := make(chan int, 1)

	if n < 0 {
		return 0, errors.New("n must be non-negative")
	}

	go func() {

		if n == 0 {
			c <- 0
			return
		}
		n1, n2 := 1, 1
		if n <= 2 {
			c <- n1
			return
		}
		var res int
		for i := 0; i < n-2; i++ {
			newN := n1 + n2
			n1 = n2
			n2 = newN
			res = newN
		}
		c <- res
	}()

	select {
	case res := <-c:
		return res, nil
	case <-time.After(timeout):
		return 0, errors.New("timeout")
	}
}

func main() {
	val, err := TimeoutFibonacci(1000, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}
