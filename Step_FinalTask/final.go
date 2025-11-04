package main

import (
	"context"
	"encoding/json"
	"io"
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

func GetTasksFromText(text string, user *string, status *string) []Ticket {
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

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	buffer := make([]byte, 1024)
	var data []byte

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			n, err := r.Read(buffer)
			if err == io.EOF {
				tmp := GetTasksFromText(string(data), user, status)

				jsonData, err := json.MarshalIndent(tmp, "", "  ")
				if err != nil {
					return err
				}

				_, err = w.Write(jsonData)
				if err != nil {
					return err
				}
				return nil
			}
			if err != nil {
				return err
			}
			data = append(data, buffer[:n]...)
		}
	}
}

//func main() {
//	r, err := os.Open("Step_FinalTask/tickets.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	w, err := os.Create("Step_FinalTask/tickets2.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer r.Close()
//
//	user := "Паша Попов"
//	status := "В работе"
//	ctx := context.Background()
//	err = GetTasks(ctx, r, w, &user, &status, time.Second*1)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
