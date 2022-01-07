package main

import (
	"cin"
	"fmt"
	"net/http"
)

func main() {
	c := cin.New()
	c.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "hello world")
	})
	c.Run(":9099")
}
