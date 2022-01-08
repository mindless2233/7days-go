package main

import (
	"cin1"
	"net/http"
)

func main() {
	r := cin.New()
	r.GET("/", func(c *cin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *cin.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *cin.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *cin.Context) {
		c.JSON(http.StatusOK, cin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
