package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GCD(a, b int) int {
	for b != 0 {
		tmp := b
		b = a % b
		a = tmp
	}
	return a
}

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])
	switch op {
	case "add":
		fmt.Fprintf(w, "%d + %d = %d", num1, num2, num1+num2)
		//fmt.Fprint(w, num1+num2)
	case "sub":
		fmt.Fprintf(w, "%d - %d = %d", num1, num2, num1-num2)
		//fmt.Fprint(w, num1-num2)
	case "mul":
		fmt.Fprintf(w, "%d * %d = %d", num1, num2, num1*num2)
		//fmt.Fprint(w, num1*num2)
	case "div":
		fmt.Fprintf(w, "%d / %d = %d", num1, num2, num1/num2)
		//fmt.Fprint(w, num1/num2)
	case "gcd":
		fmt.Fprintf(w, "gcd of %d and %d is %d", num1, num2, GCD(num1,num2))
		fmt.Fprint(w, GCD(num1,num2))
	default:
		fmt.Fprint(w, "please specify operator")
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
