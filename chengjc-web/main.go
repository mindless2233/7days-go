package main

import (
	"cin"
	"net/http"
)

func main() {
	r := cin.Default()
	r.GET("/", func(c *cin.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *cin.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
