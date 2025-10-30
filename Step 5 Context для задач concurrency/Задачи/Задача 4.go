package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func StartServer(maxTimeout time.Duration) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	readSourceHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url := "http://localhost:8081/provideData"
		res, err := http.Get(url)
		if err != nil {
			http.Error(w, "Request timeout", http.StatusServiceUnavailable)
			return
		}
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		w.WriteHeader(res.StatusCode)
		w.Write(body)
	})

	timeoutHandler := http.TimeoutHandler(
		readSourceHandler,
		maxTimeout,
		"Сервер не ответил вовремя (timeout)",
	)
	http.Handle("/readSource", timeoutHandler)
	http.ListenAndServe(":8080", nil)
}

var sleepTime time.Duration

func longHanlder(w http.ResponseWriter, r *http.Request) {
	time.Sleep(sleepTime)
	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	go func() {
		http.HandleFunc("/provideData", longHanlder)

		err := http.ListenAndServe(":8081", nil)
		if err != nil {
			fmt.Errorf("error when starting a server")
		}
	}()
	go StartServer(1 * time.Second)

	time.Sleep(100 * time.Minute)
}
