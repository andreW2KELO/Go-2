package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// В операционной системе os.Stdin похож на файл, но длина его может быть какой угодно.

	stdinScanner := bufio.NewScanner(os.Stdin)
	for stdinScanner.Scan() {
		fmt.Println("Got: ", stdinScanner.Text())
	}

	// Более того, os.Stdin в Linux — обычный файл. В стандартной библиотеке он читается так:
	// Stdin = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
}
