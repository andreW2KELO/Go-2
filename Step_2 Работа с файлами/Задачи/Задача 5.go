package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

func ExtractLog(inputFileName string, start, end time.Time) ([]string, error) {
	file, err := os.Open(inputFileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(file)
	var res []string
	for fileScanner.Scan() {
		line := fileScanner.Text()
		splittedLine := strings.Split(line, " ")
		currentTime, err := time.Parse("02.01.2006", splittedLine[0])
		if err != nil {
			return nil, err
		}
		if (currentTime.After(start) || currentTime.Equal(start)) && (currentTime.Before(end) || currentTime.Equal(end)) {
			res = append(res, line)
		}
	}
	if len(res) == 0 {
		return nil, errors.New("no logs found")
	}
	return res, nil
}

func main() {
	//start1 := time.Date(2022, 12, 19, 0, 0, 0, 0, time.UTC)
	//end2 := time.Date(2022, 12, 20, 0, 0, 0, 0, time.UTC)
	//lst, err := ExtractLog("myfile_2.txt", start1, end2)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, l := range lst {
	//	fmt.Println(l)
	//}
}
