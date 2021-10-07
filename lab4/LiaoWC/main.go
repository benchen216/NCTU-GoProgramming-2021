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
	//fmt.Fprintf(w, "hello world!")

	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	total_len := len(pathParts)
	fmt.Println(pathParts)
	fmt.Println(len(pathParts))

	if len(pathParts) < 3 {
		fmt.Fprintf(w, "hello world!")
	} else if len(pathParts) >= 3 {
		operation := pathParts[total_len-3]
		a, _ := strconv.Atoi(pathParts[total_len-2])
		b, _ := strconv.Atoi(pathParts[total_len-1])
		if operation == "add" {
			fmt.Fprintf(w, "%d + %d = %d", a, b, a+b)
		} else if operation == "sub" {
			fmt.Fprintf(w, "%d - %d = %d", a, b, a-b)
		} else if operation == "mul" {
			fmt.Fprintf(w, "%d * %d = %d", a, b, a*b)
		} else if operation == "div" {
			fmt.Fprintf(w, "%d / %d = %d, remainder = %d", a, b, a/b, a%b)
		}
	}
}

//func add(a, b int) int {
//
//}
//
//func div(a, b int) int int {
//
//}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
