package main

import (
	"sync"
	"testing"
)

func TestIncrementOnce(t *testing.T) {
	// сбросим счётчик
	counter = 0

	// функция для увеличения значения счётчика
	incrFunc = func() {
		counter++
	}

	// для ожидания горутин
	var wg sync.WaitGroup

	//  // будем делать инкремент 1000 раз
	numIncrements := 1000

	// для ожидания всех запущенных горутин
	wg.Add(numIncrements)

	// увеличиваем значение счётчика конкурентно с помощью sync.Once
	for i := 0; i < numIncrements; i++ {
		go func() {
			defer wg.Done()
			Increment()
		}()
	}

	// подождём все горутины
	wg.Wait()

	// проверим, получили ли ожидаемое значение
	expectedCounter := 1 // 1 — потому что использовали sync.Once
	actualCounter := GetCounter()

	if actualCounter != expectedCounter {
		t.Errorf("Expected counter value: %d, Actual counter value: %d", expectedCounter, actualCounter)
	}
}
