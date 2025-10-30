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

func Compare(name1, name2 string) (string, error) {
	markName1, err := getMarkByName(name1)
	if err != nil {
		return "", err
	}
	markName2, err := getMarkByName(name2)
	if err != nil {
		return "", err
	}

	switch {
	case markName1 > markName2:
		return ">", nil
	case markName1 < markName2:
		return "<", nil
	default:
		return "=", nil
	}
}
