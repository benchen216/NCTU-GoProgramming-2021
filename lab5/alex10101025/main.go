package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
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
	case "sub":
		fmt.Fprintf(w, "%d - %d = %d", num1, num2, num1-num2)
	case "mul":
		fmt.Fprintf(w, "%d * %d = %d", num1, num2, num1*num2)
	case "div":
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", num1, num2, int(num1/num2), num1%num2)
	case "gcd":
		fmt.Fprintf(w, "%s of %d and %d is %d", op, num1, num2, gcd(num1, num2))
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
