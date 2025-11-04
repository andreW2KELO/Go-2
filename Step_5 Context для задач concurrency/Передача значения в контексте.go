package main

import (
	"context"
	"fmt"
)

type userID string

func ProcessRequest(id userID) {
	// сохраним значение в контексте
	ctx := context.WithValue(context.Background(), "userID", id)
	// функция обработки
	HandleResponse(ctx)
}

// здесь контекст уже содержит userID
func HandleResponse(ctx context.Context) {
	id, ok := ctx.Value("userID").(userID)
	if !ok {
		fmt.Println("Ошибка! тип не userID")
	} else {
		fmt.Printf("handling response for (%v)", id)
	}
}

func main() {
	ProcessRequest(userID("1"))
}
