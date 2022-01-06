package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":9999", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("1")
}
