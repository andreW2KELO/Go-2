package main

import "fmt"

// Дженерики позволяют писать код, который сможет работать с различными типами данных.
// Дублировать его для каждого типа не придётся. Это упрощает и ускоряет процесс разработки.

// Определяем constraint "Number"
type Number interface {
	int | int64 | float64 | float32
}

func sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum([]int{1, 2, 3}))
}
