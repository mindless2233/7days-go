package main

import (
	"cin1"
	"fmt"
)

func main() {
	c := cin.New()

	c.GET("/", indexHandle)

	c.Run(":9099")
}

func indexHandle(c *cin.Context) {
	fmt.Fprintf(c.Writer, "sssss")
}
