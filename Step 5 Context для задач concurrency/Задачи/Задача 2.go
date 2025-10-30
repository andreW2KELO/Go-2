package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func fetchAPI(ctx context.Context, url string, timeout time.Duration) (*APIResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, context.DeadlineExceeded
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &APIResponse{Data: string(body), StatusCode: resp.StatusCode}, nil
}

type APIResponse struct {
	Data       string // тело ответа
	StatusCode int    // код ответа
}

func main() {
	//ctx := context.Background()
	//
	//url := "https://httpbin.org/delay/5"
	//timeout := 1 * time.Second
	//
	//resp, err := fetchAPI(ctx, url, timeout)
	//if err != nil {
	//	fmt.Println("Ошибка:", err)
	//	return
	//}
	//
	//fmt.Println("Код ответа:", resp.StatusCode)
	//fmt.Println("Тело ответа:", len(resp.Data))
}
