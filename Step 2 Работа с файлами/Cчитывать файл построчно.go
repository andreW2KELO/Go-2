package main

import (
	"bufio"
	"fmt"
	"os"
)

func process(line string) {
	fmt.Println(line)
}

func main() {
	f, err := os.Open("Step 2 Работа с файлами/literature.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(f) //NewScanner возвращает специальный сканер для чтения построчно
	for fileScanner.Scan() {           // Сканер перемещается к следующему токену, то есть к следующей части текста до символа новой строки
		process(fileScanner.Text()) // Функция для обработки строки с литературой — например, здесь можно попросить YaGPT дать пересказ произведения
	}
}
