package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GeneratePrimeNumbers(stop chan struct{}, prime_nums chan int, N int) {
	defer close(prime_nums) // Закрываем канал при завершении

	// Останавливаем генерацию через 0.1 секунды
	time.AfterFunc(100*time.Millisecond, func() {
		close(stop)
	})

	for i := 0; i <= N; i++ {
		select {
		case <-stop:
			// Получен сигнал остановки
			return
		default:
			if isPrime(i) {
				prime_nums <- i
			}
		}
	}
}

func main() {
	stop := make(chan struct{})
	prime_nums := make(chan int)

	go GeneratePrimeNumbers(stop, prime_nums, 10)

	for p := range prime_nums {
		fmt.Println(p)
	}
	fmt.Println("Генерация остановлена.")
}
