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
	// fmt.Fprintf(w, "hello world!")

	pathParts := strings.SplitN(r.URL.Path, "/", -1)

	url_len := len(pathParts)

	if url_len != 4 {
		fmt.Fprintf(w, "hello world!")
		return
	}

	operation := pathParts[1]
	a, _ := strconv.Atoi(pathParts[2])
	b, _ := strconv.Atoi(pathParts[3])

	if operation == "add" {
		fmt.Fprintf(w, "%d + %d = %d", a, b, add(a, b))
	} else if operation == "sub" {
		fmt.Fprintf(w, "%d - %d = %d", a, b, sub(a, b))
	} else if operation == "mul" {
		fmt.Fprintf(w, "%d * %d = %d", a, b, mul(a, b))
	} else if operation == "div" {
		quotient, remainder := div(a, b)
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", a, b, quotient, remainder)
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

func div(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
