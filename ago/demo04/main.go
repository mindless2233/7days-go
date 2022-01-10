package main

import (
	"fmt"
	"strings"
)

type Tire struct {
	pattrtn  string
	part     string
	isWild   bool
	children []*Tire
}

func (t *Tire) new() *Tire {
	return &Tire{}
}

func (t *Tire) insert(pattrtn string) {

	parts := strings.Split(pattrtn, "/")
	for _, part := range parts {
		child := t.search(part)
		if child == nil {
			child = &Tire{
				pattrtn:  part,
				isWild:   false,
				children: nil,
			}
		}
	}

}

func (t *Tire) search(pattrtn string) *Tire {
	for _, child := range t.children {
		if child.pattrtn == pattrtn {
			fmt.Printf("%s存在", pattrtn)
			return child
		}
	}
	fmt.Printf("%s不存在", pattrtn)
	return nil
}
func main() {
	root := new(Tire)
	root.insert("/cjc/123")
}
