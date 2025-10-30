package main

import (
	"bytes"
	"fmt"
	"io"
)

func Copy(r io.Reader, w io.Writer, n uint) error {
	data := io.LimitReader(r, int64(n))

	_, err := io.Copy(w, data)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	src := bytes.NewBufferString("Hello, world!")
	dst := &bytes.Buffer{}

	err := Copy(src, dst, 100)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Println(dst.String())
}
