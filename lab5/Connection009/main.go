package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)


func GCD(a, b int) int {
	if a % b == 0 { return b }

	return GCD(a, a % b)
}


func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])

	if op == "add" {
		fmt.Fprintf(w, "%d + %d = %d", num1, num2, num1+num2)
	} else if op == "sub" {
		fmt.Fprintf(w, "%d - %d = %d", num1, num2, num1-num2)
	} else if op == "mul" {
		fmt.Fprintf(w, "%d * %d = %d", num1, num2, num1*num2)
	} else if op == "div" {
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", num1, num2, num1/num2, num1%num2)
	} else if op == "gcd" {
		fmt.Fprintf(w, "gcd of %d and %d is %d", num1, num2, GCD(num1,num2) ) //gcd of 12 and 24 is 12
	} else {
		fmt.Fprintf(w, "%s", r.URL.Path)
	}
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
