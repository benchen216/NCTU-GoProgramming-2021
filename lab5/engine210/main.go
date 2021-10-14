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
		fmt.Fprint(w, num1+num2)
	case "sub":
		fmt.Fprint(w, num1-num2)
	case "mul":
		fmt.Fprint(w, num1*num2)
	case "div":
		fmt.Fprint(w, num1/num2)
	case "gcd":
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
