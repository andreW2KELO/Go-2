package main

import (
	"sync"
	"testing"
)

func TestIncrement(t *testing.T) {
	// сбросим счётчик
	counter = 0

	// для ожидания горутин
	var wg sync.WaitGroup

	// будем делать инкремент 1000 раз
	numIncrements := 1000

	// для ожидания всех запущенных горутин
	wg.Add(numIncrements)

	// увеличиваем значение счетчика конкурентно
	for i := 0; i < numIncrements; i++ {
		go func() {
			defer wg.Done()
			Increment()
		}()
	}

	// подождём все горутины
	wg.Wait()

	// проверим, получили ли ожидаемое значение
	expectedCounter := numIncrements
	actualCounter := GetCounter()

	if actualCounter != expectedCounter {
		t.Errorf("Expected counter value: %d, Actual counter value: %d", expectedCounter, actualCounter)
	}
}
