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
	if len(pathParts) != 4 {
		fmt.Fprintf(w, "hello world!")	
		return 
	}

	operation := pathParts[1]
	a, error := strconv.Atoi(pathParts[2])
	b, error := strconv.Atoi(pathParts[3])

	if error != nil {
		fmt.Fprintf(w, "hello world!")	
		return 
	}
		

	if operation == "add" {
		fmt.Fprintf(w, "%d + %d = %d\n", a , b , add(a,b) )
	} else if operation == "sub" {
		fmt.Fprintf(w, "%d - %d = %d\n", a , b , sub(a,b) )
	} else if operation == "mul" {
		fmt.Fprintf(w, "%d * %d = %d\n", a , b , mul(a,b) )
	} else if operation == "div" {
		ans1 ,ans2 := div(a,b)
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d\n", a , b , ans1 , ans2 )
	} else {
		fmt.Fprintf(w, "hello world!")
	}
	return
}

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
	return a/b , a%b
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}