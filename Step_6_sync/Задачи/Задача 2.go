package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func getUrl(name string) string {
	return fmt.Sprintf("http://localhost:8082/mark?name=%s", name)
}

func getMarkByName(name string) (int, error) {
	resp, err := http.Get(getUrl(name))
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	mark, err := strconv.Atoi(string(body))
	if err != nil {
		return 0, err
	}
	return mark, nil
}

func Average(names []string) (int, error) {
	average := 0
	for _, name := range names {
		mark, err := getMarkByName(name)
		if err != nil {
			return 0, err
		}
		average += mark
	}
	return average / len(names), nil
}
