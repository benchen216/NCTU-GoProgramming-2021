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
	return a+b
}

func sub(a, b int) int {
	return a-b
}

func mul(a, b int) int {
	return a*b
}

func div(a, b int) (int, int) {
	return a/b, a%b
}

func hello(w http.ResponseWriter, r *http.Request) {
	var r1, r2 int
	
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	
	if len(pathParts)<=3 {
		fmt.Fprintf(w, "hello world!")
		return
	}

	operation := pathParts[1]
	a, _ := strconv.Atoi(pathParts[2])
	b, _ := strconv.Atoi(pathParts[3])

	if operation == "add" {
		r1=add(a, b)
		fmt.Fprintf(w, "%d + %d = %d", a, b, r1)
	} else if operation == "sub" {
		r1=sub(a, b)
		fmt.Fprintf(w, "%d - %d = %d", a, b, r1)
	} else if operation == "mul" {
		r1=mul(a, b)
		fmt.Fprintf(w, "%d * %d = %d", a, b, r1)
	} else if operation == "div" {
		r1, r2 = div(a, b)
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", a, b, r1, r2)
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
