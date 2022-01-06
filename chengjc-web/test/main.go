package main

import (
	"cin"
	"log"
	"net/http"
	"time"
)

//func main() {
//
//	c := cin.New()
//	c.GET("/", func(c *cin.Context) {
//		c.HTML(http.StatusOK, "<h1>Hello Cin</h1>")
//	})
//	c.Run(":8123")
//
//}

//func main() {
//	r := cin.New()
//	r.GET("/", func(c *cin.Context) {
//		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//	})
//
//	r.GET("/hello", func(c *cin.Context) {
//		// expect /hello?name=geektutu
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//	})
//
//	r.GET("/hello/:name", func(c *cin.Context) {
//		// expect /hello/geektutu
//		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//	})
//
//	r.GET("/assets/*filepath", func(c *cin.Context) {
//		c.JSON(http.StatusOK, cin.H{"filepath": c.Param("filepath")})
//	})
//
//	r.Run(":9999")

//func main() {
//	r := cin.New()
//	r.GET("/index", func(c *cin.Context) {
//		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
//	})
//	v1 := r.Group("/v1")
//	{
//		v1.GET("/", func(c *cin.Context) {
//			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
//		})
//
//		v1.GET("/hello", func(c *cin.Context) {
//			// expect /hello?name=geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
//		})
//	}
//	v2 := r.Group("/v2")
//	{
//		v2.GET("/hello/:name", func(c *cin.Context) {
//			// expect /hello/geektutu
//			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
//		})
//		v2.POST("/login", func(c *cin.Context) {
//			c.JSON(http.StatusOK, cin.H{
//				"username": c.PostForm("username"),
//				"password": c.PostForm("password"),
//			})
//		})
//
//	}
//
//	r.Run(":9999")

func onlyForV2() cin.HandlerFunc {
	return func(c *cin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := cin.New()
	r.Use(cin.Logger()) // global midlleware
	r.GET("/", func(c *cin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Cin</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(context *cin.Context) {
			context.String(http.StatusOK, "%s----%s\n", context.Param("name"), context.Path)
		})
	}

	r.Run(":9012")

}
