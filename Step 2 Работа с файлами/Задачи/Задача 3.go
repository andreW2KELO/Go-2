package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFilePart(inputFilename, outFileName string, startpos int) error {
	inputFile, err := os.Open(inputFilename)
	defer inputFile.Close()
	if err != nil {
		return err
	}

	outputFile, err := os.Create(outFileName)
	defer outputFile.Close()
	if err != nil {
		return err
	}

	_, err = inputFile.Seek(int64(startpos), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := CopyFilePart("myfile.txt", "myfile_2.txt", 1024)
	if err != nil {
		fmt.Println(err)
	}
}
