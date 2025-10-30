package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// создаём экземпляр Once
	var once sync.Once
	wg := &sync.WaitGroup{}
	// функция, которая имитирует инициализацию ресурсов
	initializeResources := func() {
		time.Sleep(time.Second)
		fmt.Println("Only once initialize something")
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// в каждой горутине вызываем функцию Do
			once.Do(initializeResources)
		}()
	}
	wg.Wait()
}

// Внимательный студент мог заметить, что это похоже на паттерн Singleton.
// Использовать его нужно осознанно, так как он представляет глобальное состояние.
// Возможно, инициализация ресурсов может быть выполнена синхронно в начале работы программы.
