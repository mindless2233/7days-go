package main

import (
	"cin"
	"fmt"
	"net/http"
)

func main() {
	c := cin.New()
	c.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "根目录")
	})
	c.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q]=%q\n", k, v)
		}
	})
	c.Run(":8989")
}
