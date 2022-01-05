package main

import (
	"fmt"
	"net/http"
)

func main() {
	stu := Student{}
	stu.sayHi()
	http.ListenAndServe("1111", nil)
}

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

func (s *Student) sayHi() {

	fmt.Println("hi")
}
