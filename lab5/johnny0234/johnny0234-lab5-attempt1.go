package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])

	if op == "gcd" {
		fmt.Fprintf(w, "gcd of %d and %d is %d", num1, num2, gcd(num1, num2))
	} else if op == "add" {
		fmt.Fprintf(w, "%d + %d = %d", num1, num2, add(num1, num2))
	} else if op == "sub" {
		fmt.Fprintf(w, "%d - %d = %d", num1, num2, sub(num1, num2))
	} else if op == "mul" {
		fmt.Fprintf(w, "%d * %d = %d", num1, num2, mul(num1, num2))
	} else if op == "div" {
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", num1, num2, div(num1, num2), rem(num1, num2))
	} else {
		fmt.Fprintf(w, "hello world!")
	}
}
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func rem(a, b int) int {
	return a % b
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/Lab5", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
