package main

import (
	"fmt"
	"sync"
	"time"
)

// Этот пример извлекает несколько URL одновременно,
// используя WaitGroup для блокировки до тех пор,
// пока все извлечения не будут завершены.
type httpPkg struct{}

func (httpPkg) Get(url string) {
	time.Sleep(1 * time.Second)
	fmt.Println(url)
}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.exmple.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		wg.Go(func() {
			// Fetch the URL.
			http.Get(url)
		})
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}
