package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Давайте рассмотрим такую задачу, есть три горутины: G1, G2, G3
// горутина G1 получает данные из некоторого источника
// горутины G2 и G3 должны дождаться того, как горутина G1 получит эти данные, и что-то с ними сделать
// То-есть G2 и G3 должны быть заблокированы до того момента, пока G1 не даст им какой-то сигнал.

// Стандартная библиотека предоставляет нам объект sync.Cond.
// Конструктор его принимает в качестве аргумента
// Locker — мы будем использовать объект Mutex:

// У него есть три важные проассоциированные функции:
//
// Signal() — отправляет сигнал одной горутине
// Broadcast() — отправляет сигнал всем горутинам
// Wait() — ожидает сигнал

func listen(name string, data map[string]string, c *sync.Cond) {
	c.L.Lock()
	c.Wait()

	fmt.Printf("[%s] %s\n", name, data["key"])

	c.L.Unlock()
}

// Напишем горутину, которая получает данные и посылает слушателям сигнал о начале их обработки:
func broadcast(name string, data map[string]string, c *sync.Cond) {
	time.Sleep(time.Second)

	c.L.Lock()

	data["key"] = "value"

	fmt.Printf("[%s] данные получены\n", name)

	// отправляем сигнал слушателям
	c.Broadcast()
	c.L.Unlock()
}

func main() {
	data := map[string]string{}

	cond := sync.NewCond(&sync.Mutex{})

	go listen("слушатель 1", data, cond)
	go listen("слушатель 2", data, cond)

	go broadcast("источник", data, cond)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
