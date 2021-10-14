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
	num1, err := strconv.Atoi(r.Form["num1"][0])
	if err != nil {
		fmt.Fprintf(w, "?")
		return		
	}

	num2, err := strconv.Atoi(r.Form["num2"][0])
	if err != nil {
		fmt.Fprintf(w, "?")
		return		
	}

	switch op {
		case "add":
		fmt.Fprintf(w, "%d + %d = %d\n", num1, num2, add(num1, num2))
		
		case "sub":
		fmt.Fprintf(w, "%d - %d = %d\n", num1, num2, sub(num1, num2))
		
		case "mul":
		fmt.Fprintf(w, "%d * %d = %d\n", num1, num2, mul(num1, num2))
		
		case "div":
		if num2 == 0 {
			fmt.Fprintf(w, "divided by zero")
			return			
		}

		x, y := div(num1, num2)
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d\n", num1, num2, x, y)
		
		case "gcd":
		fmt.Fprintf(w, "gcd of %d and %d is %d", num1, num2, gcd(num1, num2))
		
		default:
		fmt.Fprintf(w, "?")
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

func gcd(a, b int) (int) {
	for b != 0 {
		a, b = b, a%b
	}
	
	return a
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
