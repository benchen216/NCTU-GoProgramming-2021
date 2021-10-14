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

	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	operation := pathParts[1]
	fmt.Printf("%#+v", pathParts)
	a, _ := strconv.Atoi(pathParts[2])
	b, _ := strconv.Atoi(pathParts[3])

	if operation == "add" {
		fmt.Fprintf(w, "%d", add(a, b))
	} else if operation == "sub" {
		fmt.Fprintf(w, "%d", sub(a, b))
	} else if operation == "mul" {
		fmt.Fprintf(w, "%d", mul(a, b))
	} else if operation == "div" {
		fmt.Fprintf(w, "%d", div(a, b))
	} else {
		fmt.Fprintf(w, "hello world!")
	}
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
