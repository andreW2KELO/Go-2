package main

import (
	"fmt"
	"os"
)

// И ещё раз про большие файлы. Предположим, у вас есть текстовый файл с логами работы программы.
// Вам нужно прочесть только последние записи из него.
// Самый простой способ — прочесть файл от начала до конца,
// но оптимальнее будет применить функцию Seek. С её помощью можно читать файл не с начала,
// а с выбранного места:

func main() {
	f, err := os.Open("Step_2 Работа с файлами/literature.txt")

	defer func(f *os.File) { // Закрытие файла
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)

	if err != nil {
		fmt.Println(err)
	}
	offset := 1800
	_, _ = f.Seek(int64(offset), 0) // Сместимся от начала файла
	buffer := make([]byte, 100)
	n, err2 := f.Read(buffer) // Прочитаем в buffer с позиции 1024
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(buffer[:n]))
}
