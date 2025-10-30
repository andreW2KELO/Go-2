package main

import (
	"fmt"
	"time"
)

func timersStoppage(v time.Timer) {
	<-v.C
	fmt.Println("Второй таймер сработал")
}

func main() {
	// создание первого таймера
	timer := time.NewTimer(3 * time.Second)
	// в канал C отправляется значение, которое указывает на окончание работы таймера
	<-timer.C
	fmt.Println("Первый таймер сработал!")
	// создание второго таймера
	timer_s := time.NewTimer(time.Second)
	// создание горутины
	go timersStoppage(*timer_s)
	// удалите // из строки ниже, чтобы сработал второй таймер
	// timersStoppage((*timer_s))
	// остановка второго таймера перед срабатыванием
	stop_s := timer_s.Stop()
	if stop_s {
		fmt.Println("Второй таймер остановлен")
	}
}
