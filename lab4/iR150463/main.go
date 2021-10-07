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
	pathParts := strings.SplitN(r.URL.Path, "/", -1)

	operation := pathParts[1]
	a, _ := strconv.Atoi(pathParts[2])
	b, _ := strconv.Atoi(pathParts[3])

	if operation == "add" {
		fmt.Fprintf(w, "%d + %d = %d", a, b, a+b)
	} else if operation == "sub" {
		fmt.Fprintf(w, "%d - %d = %d", a, b, a-b)
	} else if operation == "mul" {
		fmt.Fprintf(w, "%d * %d = %d", a, b, a*b)
	} else if operation == "div" {
		fmt.Fprintf(w, "%d / %d = %d", a, b, a/b)
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
