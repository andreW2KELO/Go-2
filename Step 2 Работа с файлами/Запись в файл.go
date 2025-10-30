package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("C:\\Users\\79968\\Documents\\myfile.txt.txt")
	defer func() {
		errClose := f.Close() // Не забываем закрыть и проверить на ошибку — при записи файла это важно
		if errClose != nil {
			fmt.Println(errClose)
		}
	}()
	if err != nil {
		fmt.Println(err, "wefewwe")
	}
	// Если файл существует, он будет перезаписан, а если не существует — создан.
	// Теперь мы можем записать данные с помощью дескриптора:
	n, err2 := f.WriteString("hello") // n — сколько байтов записали
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("Байт записано:", n)
	// А если нам нужно дописать данные в уже существующий файл?
	// Функция os.Create заменит его на пустой. Чтобы этого не случилось, воспользуемся os.OpenFile:
}
