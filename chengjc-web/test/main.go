package main

import (
	"cin"
	"net/http"
)

func main() {

	c := cin.New()
	c.GET("/", func(c *cin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Cin</h1>")
	})
	c.Run(":8123")

}
