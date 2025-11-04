package main

import (
	"fmt"
	"reflect"
)

// Чтобы получить значение поля структуры, будет полезен метод FieldByName объекта reflect.Value.
// Этот метод принимает имя поля в качестве аргумента и возвращает reflect.Value со значением поля:

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}
	v := reflect.ValueOf(p)
	name := v.FieldByName("Name").String()
	age := v.FieldByName("Age").Int()
	fmt.Println(name, age)
}

// Здесь мы создаём структуру Person с полями Name и Age.
// Затем получаем объект reflect.Value для структуры p и используем метод FieldByName,
// чтобы получить значения полей Name и Age. Далее преобразуем значение поля Name в строку с помощью метода String,
// а значение поля Age — в число с помощью метода Int.
