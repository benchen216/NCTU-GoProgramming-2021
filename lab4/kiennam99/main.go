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
		
	 	operation := pathParts[len(pathParts)-3]
	 	a, _ := strconv.Atoi(pathParts[len(pathParts)-2])
	 	b, _ := strconv.Atoi(pathParts[len(pathParts)-1])
	
	 	if operation == "add" {
			fmt.Fprintf(w, add(a,b))

	 	} else if operation == "sub" {
			fmt.Fprintf(w, sub(a,b))

	 	} else if operation == "mul" {
			fmt.Fprintf(w, mul(a,b))

	 	} else if operation == "div" {
			
			fmt.Fprintf(w, div(a,b))

	 	} else {
	 		fmt.Fprintf(w, "hello world!")
	 	}
}

func add(a, b int) string {
	output := strconv.Itoa(a) + " + " + strconv.Itoa(b) + " = "
	car := a+b
	output += strconv.Itoa(car)
	
	return output
}

func sub(a, b int) string {
	output := strconv.Itoa(a) + " - " + strconv.Itoa(b) + " = "
	car := a-b
	output += strconv.Itoa(car)

	return output
}
func mul(a, b int) string {
	output := strconv.Itoa(a) + " * " + strconv.Itoa(b) + " = "
	car := a*b
	output += strconv.Itoa(car)

	return output
}
func div(a, b int) string {
	output := strconv.Itoa(a) + " / " + strconv.Itoa(b) + " = "
	car := a/b
	output += strconv.Itoa(car) + ", remainder = " + strconv.Itoa(a-a/b*b)

	return output
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
