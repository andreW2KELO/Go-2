package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadString(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	reader := strings.NewReader("Привет, Арсений!")
	text, err := ReadString(reader)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(text)
}
