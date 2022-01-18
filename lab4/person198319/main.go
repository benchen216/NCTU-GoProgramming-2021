package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")

	pathParts := strings.SplitN(r.URL.Path, "/", -1)

	operation := pathParts[0]
	a, _ := strconv.Atoi(pathParts[1])
	b, _ := strconv.Atoi(pathParts[2])
	var c int
	if operation == "add" {
		c = add(a, b)
		x := "+"
		y := "="
		fmt.Fprint(w, a, x, b, y, c)
	} else if operation == "sub" {
		c = sub(a, b)
		x := "-"
		y := "="
		fmt.Fprint(w, a, x, b, y, c)
	} else if operation == "mul" {
		c = mul(a, b)
		x := "*"
		y := "="
		fmt.Fprint(w, a, x, b, y, c)
	} else if operation == "div" {
		c = div(a, b)
		x := "/"
		y := "="
		fmt.Fprint(w, a, x, b, y, c)
	} else {
		fmt.Fprintf(w, "hello world!")
	}
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
