package main

import (
	"bytes"
	"context"
	"io"
)

func Contains(ctx context.Context, r io.Reader, seq []byte) (bool, error) {
	bufSize := 4096
	buffer := make([]byte, bufSize)
	var tail []byte

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		default:
			n, err := r.Read(buffer)
			if n > 0 {

				chunk := append(tail, buffer[:n]...)

				if bytes.Contains(chunk, seq) {
					return true, nil
				}

				if len(chunk) >= len(seq)-1 {
					tail = chunk[len(chunk)-len(seq)+1:]
				} else {
					tail = chunk
				}
			}

			if err == io.EOF {
				return false, nil
			}
			if err != nil {
				return false, err
			}
		}
	}
}

func main() {
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	//
	//data := strings.NewReader("Hello, world! This is some test data.")
	//found, err := Contains(ctx, data, []byte("test"))
	//
	//fmt.Println(found, err)
}
