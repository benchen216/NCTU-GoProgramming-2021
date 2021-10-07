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

	if (len(pathParts) >= 4) {
		operation := pathParts[1]
		a, _ := strconv.Atoi(pathParts[2])
		b, _ := strconv.Atoi(pathParts[3])
	
		if operation == "add" {
			fmt.Fprintf(w, strconv.Itoa(a)+" + "+strconv.Itoa(b)+" = "+strconv.Itoa(a+b))
		} else if operation == "sub" {
			fmt.Fprintf(w, strconv.Itoa(a)+" - "+strconv.Itoa(b)+" = "+strconv.Itoa(a-b))
		} else if operation == "mul" {
			fmt.Fprintf(w, strconv.Itoa(a)+" * "+strconv.Itoa(b)+" = "+strconv.Itoa(a*b))
		} else if operation == "div" {
			fmt.Fprintf(w, strconv.Itoa(a)+" / "+strconv.Itoa(b)+" = "+strconv.Itoa(a/b)+", remainder = "+strconv.Itoa(a%b))
		} else {
			fmt.Fprintf(w, "hello world!")
		}
	} else {
		fmt.Fprintf(w, "hello world!")
	}
}

// func add(a, b int) int {
// 	return a + b
// }

// func div(a, b int) int int {
// 	return a / b, a % b
// }

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
