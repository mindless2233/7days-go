package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {
}

func (e Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Println("/")
	case "/hello":
		fmt.Fprintf(w, "hello\n")
	default:
		fmt.Fprintf(w, "default\n")
	}
}
func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
