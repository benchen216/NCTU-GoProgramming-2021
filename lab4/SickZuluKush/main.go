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

	if len(pathParts) != 4 {
		fmt.Fprintf(w, "hello world!")
		return
	}
	
	operation := pathParts[1]
	a, err := strconv.Atoi(pathParts[2])
	if err != nil {
		fmt.Fprintf(w, "hello world!")
		return		
	}
	
	b, err := strconv.Atoi(pathParts[3])
	if err != nil {
		fmt.Fprintf(w, "hello world!")
		return		
	}

	if operation == "add" {
		fmt.Fprintf(w, "%d + %d = %d\n", a, b, add(a, b))
	} else if operation == "sub" {
		fmt.Fprintf(w, "%d - %d = %d\n", a, b, sub(a, b))
	} else if operation == "mul" {
		fmt.Fprintf(w, "%d * %d = %d\n", a, b, mul(a, b))
	} else if operation == "div" {
		if b == 0 {
			fmt.Fprintf(w, "divided by zero")
			return			
		}
		
		x, y := div(a, b)
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d\n", a, b, x, y)
	} else {
		fmt.Fprintf(w, "hello world!")
	}
}

func add(a, b int) (int) {
	return a + b
}

func sub(a, b int) (int) {
	return a - b
}

func mul(a, b int) (int) {
	return a * b
}

func div(a, b int) (int , int) {
	return a / b, a % b
}

func main() {
	port := "8500"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}