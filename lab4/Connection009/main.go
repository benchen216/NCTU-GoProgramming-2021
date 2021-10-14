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

	len_url   := len(pathParts)
	operation := pathParts[len_url-3]
	a, _ 	  := strconv.Atoi(pathParts[len_url-2])
	b, _ 	  := strconv.Atoi(pathParts[len_url-1])

	if operation == "add" {
		fmt.Fprintf(w, "%d + %d = %d", a, b, a+b)
	} else if operation == "sub" {
		fmt.Fprintf(w, "%d - %d = %d", a, b, a-b)
	} else if operation == "mul" {
		fmt.Fprintf(w, "%d * %d = %d", a, b, a*b)
	} else if operation == "div" {
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", a, b, a/b, a%b)
	} else {
		fmt.Fprintf(w, "%s", r.URL.Path)
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