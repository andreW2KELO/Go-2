package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func FetchAPI(ctx context.Context, urls []string, timeout time.Duration) []*APIResponse {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var wg sync.WaitGroup
	var result = make([]*APIResponse, len(urls))

	wg.Add(len(urls))
	for i, url := range urls {
		go func(i int, url string) {
			defer wg.Done()

			exmplApiResp := &APIResponse{URL: url}

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				exmplApiResp.Err = err
				result[i] = exmplApiResp
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				if ctx.Err() == context.DeadlineExceeded {
					exmplApiResp.Err = context.DeadlineExceeded
				} else {
					exmplApiResp.Err = err
				}
				result[i] = exmplApiResp
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				exmplApiResp.Err = err
				result[i] = exmplApiResp
				return
			}

			exmplApiResp.Data = string(body)
			exmplApiResp.StatusCode = resp.StatusCode
			result[i] = exmplApiResp

		}(i, url)
	}

	wg.Wait()
	return result
}

type APIResponse struct {
	URL        string // запрошенный URL
	Data       string // тело ответа
	StatusCode int    // код ответа
	Err        error  // ошибка, если возникла
}

func main() {
	ctx := context.Background()
	urls := []string{
		"https://www.twitch.tv/",
		"https://www.twitch.tv/",
		"https://www.twitch.tv/",
		"https://httpbin.org/delay/3",
		"https://example.com",
	}

	timeout := 2 * time.Second

	results := FetchAPI(ctx, urls, timeout)

	for _, r := range results {
		fmt.Println("🔹 URL:", r.URL)
		fmt.Println("   Status:", r.StatusCode)
		fmt.Println("   Error:", r.Err)
		if len(r.Data) > 50 {
			fmt.Println("   Data:", r.Data[:50]+"...")
		} else {
			fmt.Println("   Data:", r.Data)
		}
		fmt.Println()
	}
}
