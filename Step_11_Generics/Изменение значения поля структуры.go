package main

import (
	"fmt"
	"reflect"
)

// Чтобы изменить значения поля структуры, можно использовать метод SetField объекта reflect.Value.
// Этот метод принимает имя поля и новое значение в качестве аргументов:

type Person2 struct {
	Name string
	Age  int
}

func main() {
	p := Person2{Name: "John", Age: 30}
	v := reflect.ValueOf(&p).Elem()
	v.FieldByName("Name").SetString("Jane")
	v.FieldByName("Age").SetInt(25)
	fmt.Println(p)
}
