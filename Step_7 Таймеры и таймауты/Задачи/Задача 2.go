package main

import (
	"fmt"
	"strings"
	"time"
)

func QuizRunner(questions, answers []string, answerCh chan string) int {
	correct := 0

	for i := range questions {
		timeout := make(chan bool, 1)
		go func() {
			time.Sleep(1 * time.Second)
			timeout <- true
		}()

		select {
		case <-timeout:
			continue
		case answer := <-answerCh:
			userAnswer := strings.TrimSpace(strings.ToLower(answer))
			correctAnswer := strings.TrimSpace(strings.ToLower(answers[i]))

			if userAnswer == correctAnswer {
				correct++
			}
		}
	}

	return correct
}

// Пример использования:
func main() {
	questions := []string{"Столица Франции?", "2 + 2?", "Язык Go разработан компанией...?"}
	answers := []string{"Париж", "4", "Google"}

	answerCh := make(chan string)

	// Симулируем ответы пользователя в отдельной горутине
	go func() {
		time.Sleep(100 * time.Millisecond)
		answerCh <- "париЖ" // правильный (регистр не важен)
		time.Sleep(1100 * time.Millisecond)
		answerCh <- "4" // неправильный, но опоздал
		time.Sleep(50 * time.Millisecond)
		answerCh <- "google" // правильный
	}()

	score := QuizRunner(questions, answers, answerCh)
	fmt.Println("Количество правильных ответов:", score)
}
