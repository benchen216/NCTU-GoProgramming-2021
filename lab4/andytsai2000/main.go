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

	operation := pathParts[0]
	a, _ := strconv.Atoi(pathParts[1])
	b, _ := strconv.Atoi(pathParts[2])

	if operation == "add" {
		s := strconv.Itoa(a) + " + " + strconv.Itoa(b) + " = " + strconv.Itoa(a+b)
		fmt.Fprint(w, s)
	} else if operation == "sub" {
		s := strconv.Itoa(a) + " - " + strconv.Itoa(b) + " = " + strconv.Itoa(a-b)
		fmt.Fprint(w, s)
	} else if operation == "mul" {
		s := strconv.Itoa(a) + " * " + strconv.Itoa(b) + " = " + strconv.Itoa(a*b)
		fmt.Fprint(w, s)
	} else if operation == "div" {
		s := strconv.Itoa(a) + " / " + strconv.Itoa(b) + " = " + strconv.Itoa(a/b)
		s += ", remainder = " + strconv.Itoa(a%b)
		fmt.Fprint(w, s)
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
