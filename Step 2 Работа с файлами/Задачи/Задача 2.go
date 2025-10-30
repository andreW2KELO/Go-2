package main

import (
	"bufio"
	"os"
)

func LineByNum(inputFilename string, lineNum int) string {
	f, err := os.Open(inputFilename)
	if err != nil {
		return ""
	}
	fileScanner := bufio.NewScanner(f)

	current := 0

	for fileScanner.Scan() {
		if lineNum == current {
			return fileScanner.Text()
		}
		current++
	}
	return ""
}
