package main

import "fmt"

func myFunc(a interface{}) {
	s, ok := a.(string)
	if ok {
		fmt.Printf("'%v' is a string\n", s)
	} else {
		fmt.Printf("'%v' is not a string\n", a)
	}
}

func main() {
	myFunc("hello")
	myFunc(42)
}
