package main

import (
	"sync"
	"testing"
)

var numGoroutines = 10

func TestConcurrentProcessing(t *testing.T) {
	var wg sync.WaitGroup

	// инициализация

	// запуск горутин
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// логика горутины
		}(i)
	}

	// ожидание завершения всех горутин
	wg.Wait()

	// проверки результатов
}
