package main

import (
	"fmt"
	"os"
)

func ReadContent(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(f)
}

func main() {
	fmt.Println(ReadContent("myfile.txt"))
}
