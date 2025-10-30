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

func CompareList(names []string) (map[string]string, error) {
	marks := map[string]int{}
	average := 0
	res := map[string]string{}
	for _, name := range names {
		mark, err := getMarkByName(name)
		if err != nil {
			return res, err
		}
		marks[name] = mark
		average += mark
	}

	average /= len(names)
	for name, mark := range marks {
		switch {
		case mark > average:
			res[name] = ">"
		case mark < average:
			res[name] = "<"
		default:
			res[name] = "="
		}
	}

	return res, nil
}
