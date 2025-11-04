package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func readSource(ctx context.Context) error {
	// имитируем долгую работу функции
	time.Sleep(3 * time.Second)
	// допустим, возникла ошибка в процессе
	return fmt.Errorf("some error in readSource")
}

func processSourceData(ctx context.Context) error {
	// получаем данные в цикле
	for {
		select {
		// раз в секунду получаем новые данные
		case <-time.After(time.Second):
			// здесь может быть код получения очередной порции данных
			fmt.Println("process data bit by bit...")
		// проверим контекст на отмену
		case <-ctx.Done():
			fmt.Println("processSourceData was canceled")
			return nil
		}
	}
}

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()
	// ожидаем завершения горутин
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		// запускаем функцию обработки данных
		if err := processSourceData(ctxWithCancel); err != nil {
			fmt.Printf("processSourceData(ctx): %s", err)
		}
	}()
	go func() {
		defer wg.Done()
		// запускаем функцию чтения данных
		if err := readSource(ctxWithCancel); err != nil {
			cancelCtx()
			fmt.Printf("readSource(ctx): %s", err)
		}
	}()
	// ждём завершения
	wg.Wait()
}
