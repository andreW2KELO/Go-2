package main

import (
	"fmt"
	"os"
)

const flags = `
    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
`

func appendInfoInFile(filename string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0600)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	// Флаг os.O_APPEND указывает, что файл открывается в режиме добавления.
	// Теперь при записи через дескриптор f исходные данные останутся в файле.
	// Число 0600 предоставляет права доступа к файловой системе.
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := f.WriteString(flags)
	if err2 != nil {
		fmt.Println(err2)
	}
}

func appendInfoInFileWithShift(filename string) {
	f, err := os.OpenFile(filename, os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	var offset int64 = 1024
	_, err1 := f.Seek(offset, 0)
	if err1 != nil {
		fmt.Println(err1)
	}
	_, err2 := f.WriteString("offset write")
	if err2 != nil {
		fmt.Println(err2)
	}
}

func main() {
	//appendInfoInFile("C:\\Users\\79968\\Documents\\myfile.txt.txt")
	appendInfoInFileWithShift("myfile.txt")
}
