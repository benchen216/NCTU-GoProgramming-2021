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

		if len(pathParts)<=3 {
			fmt.Fprintf(w, "hello world!")
			return
		}
		
	 	operation := pathParts[1]
	 	a, _ := strconv.Atoi(pathParts[2])
	 	b, _ := strconv.Atoi(pathParts[3])
	 	if operation == "add" {
			fmt.Fprintf(w,"%d + %d = %d",a,b,Add(a,b))

	 	} else if operation == "sub" {
			fmt.Fprintf(w,"%d - %d = %d",a,b,Sub(a,b))

	 	} else if operation == "mul" {
			fmt.Fprintf(w,"%d * %d = %d",a,b,Mul(a,b))

	 	} else if operation == "div" {
			fmt.Fprintf(w,"%d / %d = %d, remainder = %d",a,b,a/b,a%b)

	 	} else {
	 		fmt.Fprintf(w, "hello world!")
	 	}
}

func Add(x,y int) int {
	return x+y
}

func Sub(x,y int) int {
	return x-y
}

func Mul(x,y int) int {
	return x*y
}

func Div(x,y int) int {
	return x/y
}


func main() {

	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}