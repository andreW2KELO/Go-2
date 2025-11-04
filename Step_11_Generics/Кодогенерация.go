package main

import (
	"os"
	"text/template"
)

// Кодогенерация — это генерация кода на основе шаблонов и данных.

//Она пригодится в случаях, когда нужно создать много кода с одинаковой структурой и различиями в некоторых параметрах.
// С помощью кодогенерации вы можете создавать операции для базы данных или файлы конфигурации
// (например, для Kubernetes — системы разворачивания приложений в кластере).

// В Go для кодогенерации можно использовать стандартный пакет text/template.
// Он позволяет создавать шаблоны, заполнять их данными и генерировать код.

// Предположим, мы создаём CRUD-операции (Create-Read-Update-Delete, стандартный набор операций) для базы данных.
// У нас есть структура User:

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func main() {
	t := template.Must(template.ParseFiles("template.go.tmpl"))
	data := struct {
		Package string
		Struct  string
		User    string
	}{
		Package: "db",
		Struct:  "UserRepository",
		User:    "User",
	}
	t.Execute(os.Stdout, data)
}
