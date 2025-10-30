package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func readJSON(ctx context.Context, path string, result chan<- []byte) {
	defer close(result)

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	var data []byte

	for {
		select {
		case <-ctx.Done():
			return
		default:
			n, err := file.Read(buffer)
			if err != nil {
				if err == io.EOF {
					result <- data
				} else {
					return
				}
				return
			}
			data = append(data, buffer[:n]...)
		}
	}
}

func main() {
	// Создаем канал для результатов
	resultChan := make(chan []byte)

	// Создаем контекст с таймаутом 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Запускаем чтение файла
	go readJSON(ctx, "test.json", resultChan)

	// Ждем получения данных или завершения контекста
	select {
	case data := <-resultChan:
		// Парсим JSON
		fmt.Printf("Данные из JSON: %+v\n", data)
	case <-ctx.Done():
		log.Printf("Таймаут при чтении файла")
	}
}
