package main

import "fmt"

type mfc func(a, b int) int

func sum(a, b int) int {
	return a + b
}

func diff(a, b int) int {
	return a - b
}

func main() {
	var fc mfc
	fc = sum
	fmt.Println(fc(1, 2))

	fb := diff
	fmt.Println(fb(1, 2))

}
