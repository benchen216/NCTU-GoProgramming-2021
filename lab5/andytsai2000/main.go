package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func parse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Connected!\n")
	r.ParseForm()
	fmt.Fprintf(w, "The method is: %s\n", r.Method)
	fmt.Fprintf(w, "The form is: %s\n", r.Form)
	fmt.Fprintf(w, "The path is: %s\n", r.URL.Path)
	fmt.Fprintf(w, "The p3 parameter is: %s\n\n", r.Form["p3"])
	for k, v := range r.Form {
		fmt.Fprintf(w, "Parameter: %s\n", k)
		fmt.Fprintf(w, "Value: %s\n", strings.Join(v, ""))
	}
}

func cal(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	operation := r.Form.Get("op")
	a, _ := strconv.Atoi(r.Form.Get("num1"))
	b, _ := strconv.Atoi(r.Form.Get("num2"))

	if operation == "add" {
		fmt.Fprintf(w, "%d + %d = %d\n", a, b, a+b)
	} else if operation == "sub" {
		fmt.Fprintf(w, "%d - %d = %d\n", a, b, a-b)
	} else if operation == "mul" {
		fmt.Fprintf(w, "%d * %d = %d\n", a, b, a*b)
	} else if operation == "div" {
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d\n", a, b, a/b, a%b)
	} else if operation == "gcd" {
		var r int
		ra := a
		rb := b
		for rb != 0 {
			r = ra % rb
			ra = rb
			rb = r
		}
		fmt.Fprintf(w, "gcd of %d and %d is %d\n", a, b, ra)
	} else {
		fmt.Fprintf(w, "error")
	}
}

func main() {
	port := "8500"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	http.HandleFunc("/", parse)
	http.HandleFunc("/Lab5", cal)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
