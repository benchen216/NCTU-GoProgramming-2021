package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func helloWorld(w http.ResponseWriter) {
	fmt.Fprintf(w, "hello world!")
}

func Add(w http.ResponseWriter, num1, num2 int) {
	fmt.Fprintf(w, "%d + %d = %d", num1, num2, num1+num2)
}

func Sub(w http.ResponseWriter, num1, num2 int) {
	fmt.Fprintf(w, "%d - %d = %d", num1, num2, num1-num2)
}

func Mul(w http.ResponseWriter, num1, num2 int) {
	fmt.Fprintf(w, "%d * %d = %d", num1, num2, num1*num2)
}

func Div(w http.ResponseWriter, num1, num2 int) {
	fmt.Fprintf(w, "%d / %d = %d, remainder = %d", num1, num2, num1/num2, num1%num2)
}

func GCD(w http.ResponseWriter, num1, num2 int) {
	a := num1
	b := num2

	if b > a {
		a, b = b, a
	}

	for a != b {
		a = a - b
		if b > a {
			a, b = b, a
		}
	}

	fmt.Fprintf(w, "gcd of %d and %d is %d", num1, num2, a)
}

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])

	switch op {
	case "add":
		Add(w, num1, num2)
	case "sub":
		Sub(w, num1, num2)
	case "mul":
		Mul(w, num1, num2)
	case "div":
		Div(w, num1, num2)
	case "gcd":
		GCD(w, num1, num2)
	default:
		helloWorld(w)
	}
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/Lab5", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
