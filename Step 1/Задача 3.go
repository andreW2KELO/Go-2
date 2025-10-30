package main

import (
	"fmt"
	"strings"
)

type UpperWriter struct {
	UpperString string
}

func (u *UpperWriter) Write(p []byte) (n int, err error) {
	upper := strings.ToUpper(string(p))
	u.UpperString = upper

	return len(p), nil
}

func main() {
	u := UpperWriter{}

	_, _ = u.Write([]byte("hello"))
	fmt.Println(u.UpperString)

	_, _ = u.Write([]byte(" world"))
	fmt.Println(u.UpperString) // Вывод: HELLO WORLD
}
