package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")

	// 	pathParts := strings.SplitN(r.URL.Path, "/", -1)

	// 	operation := pathParts[?]
	// 	a, _ := strconv.Atoi(pathParts[?])
	// 	b, _ := strconv.Atoi(pathParts[?])

	// 	if operation == "add" {

	// 	} else if operation == "sub" {

	// 	} else if operation == "mul" {

	// 	} else if operation == "div" {

	// 	} else {
	// 		fmt.Fprintf(w, "hello world!")
	// 	}
}

func add(a, b int) int {

}

func div(a, b int) int int {

}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}