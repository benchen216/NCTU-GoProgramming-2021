package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func gcd(a, b int) int {
	max := a
	if a > b {
		max = b
	}
	var i int
	for i = max; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			break
		}
	}
	return i
}

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	a, _ := strconv.Atoi(r.Form["num1"][0])
	b, _ := strconv.Atoi(r.Form["num2"][0])
	switch op {
	case "add":
		fmt.Fprintf(w, "%d + %d = %d", a, b, add(a, b))
	case "sub":
		fmt.Fprintf(w, "%d - %d = %d", a, b, sub(a, b))
	case "mul":
		fmt.Fprintf(w, "%d * %d = %d", a, b, mul(a, b))
	case "div":
		val, remainder := div(a, b)
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", a, b, val, remainder)
	case "gcd":
		fmt.Fprintf(w, "gcd of %d and %d is %d", a, b, gcd(a, b))
	default:
		fmt.Fprintf(w, "hello world!")
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
