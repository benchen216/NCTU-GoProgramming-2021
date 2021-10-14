package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Add(a, b int) (int) {
	return (a+b)
}

func Sub(a, b int) (int) {
	return (a-b)
}

func Mul(a, b int) (int) {
	return (a*b)
}

func Div(a, b int) (int, int) {
	return (a/b), (a%b)
}

func GCD(a, b int) (int) {
	if b == 0 {
		return a
	} else {
		return GCD(b, a%b)
	}
}

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])
	switch op {
		case "add":
			fmt.Fprintf(w,"%d + %d = %d",num1, num2, Add(num1, num2))
		case "sub":
			fmt.Fprintf(w,"%d - %d = %d",num1, num2, Sub(num1, num2))
		case "mul":
			fmt.Fprintf(w,"%d * %d = %d",num1, num2, Mul(num1, num2))
		case "div":
			q, r := Div(num1, num2)
			fmt.Fprintf(w,"%d / %d = %d, remainder = %d",num1, num2, q, r)
		case "gcd":
			fmt.Fprintf(w,"gcd of %d and %d is %d",num1, num2, GCD(num1, num2))
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
