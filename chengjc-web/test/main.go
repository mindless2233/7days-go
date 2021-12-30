package main

import (
	"cin"
	"fmt"
	"net/http"
)

func main() {

	c := cin.New()
	c.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello cin!")
	})
	c.Run(":8123")

}
