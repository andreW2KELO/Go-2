package main

import (
	"fmt"
	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

var validStatuses = map[string]bool{
	"Готово":           true,
	"В работе":         true,
	"Не будет сделано": true,
}

func normalizeSpaces(s string) string {
	s = strings.TrimSpace(s)
	fields := strings.Fields(s)
	return strings.Join(fields, " ")
}

func GetTasks(text string, user *string, status *string) []Ticket {
	messages := strings.Split(strings.TrimSpace(text), "\n")
	var tickets []Ticket

	for _, message := range messages {
		message = strings.TrimSpace(message)
		if message == "" || strings.Index(message, "TICKET") != 0 {
			continue
		}

		m := strings.Split(message, "_")

		if len(m) != 4 {
			continue
		}

		ticketM := normalizeSpaces(m[0])
		nameM := normalizeSpaces(m[1])
		statusM := normalizeSpaces(m[2])
		dateM := normalizeSpaces(m[3])

		if !validStatuses[statusM] {
			continue
		}

		timeStamp, err := time.Parse("2006-01-02", dateM)

		if err != nil {
			continue
		}

		if user != nil && *user != nameM {
			continue
		}
		if status != nil && *status != statusM {
			continue
		}

		tickets = append(tickets, Ticket{
			Ticket: ticketM,
			User:   nameM,
			Status: statusM,
			Date:   timeStamp,
		})

	}
	return tickets
}

func main() {
	chatHistory := `
	TICKET-12345_Паша Попов_Готово_2024-01-01
	TICKET-12346_Иван Иванов_В работе_2024-01-02
	TICKET-12347_Анна Смирнова_Не будет сделано_2024-01-03
	TICKET-12348_Паша Попов_В работе_2024-01-04
	`
	strings.Split(strings.TrimSpace(chatHistory), "\n")
	fmt.Println(chatHistory)
	user := "Паша Попов"
	//user := "Паша Попов"
	tasks := GetTasks(chatHistory, &user, nil)
	fmt.Println(tasks)
}
