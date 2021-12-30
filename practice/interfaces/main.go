package main

import "fmt"

type Person interface {
	getName() string
}
type Student struct {
	name string
	age  int
}

func (s Student) getName() string {
	return s.name
}

func main() {
	var p Person = &Student{name: "111", age: 12}
	fmt.Println(p.getName())
}
