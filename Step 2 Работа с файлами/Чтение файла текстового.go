package main

import (
	"fmt"
	"os"
)

func readFileLiterature(filename string) (string, error) {
	data, err := os.ReadFile("Step 2 Работа с файлами/" + filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	reader, _ := readFileLiterature("literature.txt")
	fmt.Println(reader)
}
